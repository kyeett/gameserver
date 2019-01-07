package grpc

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/improbable-eng/grpc-web/go/grpcweb"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/kyeett/gameserver/entity"
	"github.com/kyeett/gameserver/localserver"
	pb "github.com/kyeett/gameserver/proto"
	"github.com/kyeett/gameserver/types"
)

var _ pb.BackendServer = (*GrpcServer)(nil)

type GrpcServer struct {
	local *localserver.LocalServer
	mu    *sync.RWMutex
}

func NewServer(w types.World) (*GrpcServer, error) {
	log.Info("new remote server created")
	l := localserver.New(w).(*localserver.LocalServer)

	return &GrpcServer{
		l,
		&sync.RWMutex{},
	}, nil
}

// Todo clean up
func (s *GrpcServer) Run(ctx context.Context, port string, secure bool) {
	if secure {
		s.RunWeb(ctx, port)
		return
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	ss := grpc.NewServer()
	pb.RegisterBackendServer(ss, s)
	// Register reflection service on gRPC server.
	reflection.Register(ss)

	go func() {
		<-ctx.Done()
		ss.GracefulStop()
	}()
	log.Infof("starting server at %s", port)
	if err := ss.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *GrpcServer) RunWeb(ctx context.Context, port string) {
	log.Infof("start a http server at %s", port)

	ss := grpc.NewServer()
	pb.RegisterBackendServer(ss, s)
	fmt.Println(s.local.World())

	wrappedServer := grpcweb.WrapServer(ss,
		grpcweb.WithWebsockets(true),
		grpcweb.WithCorsForRegisteredEndpointsOnly(false),
		grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool { return true }),
		grpcweb.WithOriginFunc(func(origin string) bool { return true }),
	)

	handler := func(resp http.ResponseWriter, req *http.Request) {

		// log.Println(req)
		log.Println("Trolo:", req.ProtoMajor,
			wrappedServer.IsAcceptableGrpcCorsRequest(req),
			websocket.IsWebSocketUpgrade(req),
			strings.Contains(req.Header.Get("Content-Type"), "application/grpc"))
		log.Println(req.URL, ss.GetServiceInfo())
		log.Println(req.ProtoMajor == 2, strings.Contains(req.Header.Get("Content-Type"), "application/grpc"),
			websocket.IsWebSocketUpgrade(req))

		log.Println()

		if req.Method == "OPTIONS" {
			allowCors(resp, req)
			return
		}

		log.Println(req.Header)

		if strings.Contains(req.Header.Get("Content-Type"), "application/grpc") || websocket.IsWebSocketUpgrade(req) {
			log.Println("In here!")
			wrappedServer.ServeHTTP(resp, req)
		} else {
			log.Println("Serve files!", req)
		}
	}

	addr := "localhost:" + port
	httpsSrv := &http.Server{
		Addr:    addr,
		Handler: http.HandlerFunc(handler),
		// Some security settings
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       120 * time.Second,
		TLSConfig: &tls.Config{
			PreferServerCipherSuites: true,
			CurvePreferences: []tls.CurveID{
				tls.CurveP256,
				tls.X25519,
			},
		},
	}

	log.Infof("Starting server on %s\n", addr)
	log.Fatal(httpsSrv.ListenAndServeTLS("../cert.pem", "../key.pem"))
	// log.Fatal(httpsSrv.ListenAndServeTLS("./cert.pem", "./key.pem"))
}

func (s *GrpcServer) NewPlayer(ctx context.Context, _ *pb.Empty) (*pb.PlayerID, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	p, err := s.local.NewPlayer()
	if err != nil {
		return nil, err
	}

	return &pb.PlayerID{ID: p.ID}, nil
}

// Todo, decide format to send over wire
func (s *GrpcServer) WorldRequest(ctx context.Context, _ *pb.Empty) (*pb.WorldResponse, error) {
	log.Info("got WorldRequest from client")
	log.Error(*s)
	log.Error(s.local)
	log.Info("got WorldRequest from client")
	return &pb.WorldResponse{
		Tiles:  s.local.World().TileBytes(),
		Width:  int32(s.local.World().Width),
		Height: int32(s.local.World().Height),
	}, nil
}
func (s *GrpcServer) EntityStream(_ *pb.Empty, stream pb.Backend_EntityStreamServer) error {

	ticker := time.NewTicker(5 * time.Millisecond)
	for {
		<-ticker.C

		s.mu.Lock()
		e := s.local.Entities()
		payload, err := GobMarshal(&e)
		s.mu.Unlock()
		if err != nil {
			return err
		}

		stream.Send(&pb.EntityResponse{
			Payload: payload,
		})

	}
}

//Todo: refactor this code.
func (s *GrpcServer) PerformAction(ctx context.Context, req *pb.ActionRequest) (*pb.ActionResponse, error) {
	tmpEntity := req.GetEntity()

	s.mu.Lock()
	defer s.mu.Unlock()
	e, err := s.local.PerformAction(entity.Entity{ID: tmpEntity.GetID()},
		types.Position{
			Coord: types.Coord{X: int(tmpEntity.GetX()), Y: int(tmpEntity.GetY())},
			Theta: int(tmpEntity.GetTheta())})
	if err != nil {
		return nil, err
	}
	ent := pb.Entity{
		ID:    e.ID,
		X:     int32(e.Position.X),
		Y:     int32(e.Position.Y),
		Theta: int32(e.Position.Theta),
	}
	return &pb.ActionResponse{Entity: &ent}, nil
}

func allowCors(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, x-grpc-web")
}
