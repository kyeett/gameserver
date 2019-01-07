// +build !js

package grpc

import (
	pb "github.com/kyeett/gameserver/proto"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var cert = "../cert.pem"

func CorrectClient(serverAddr string, secure bool) (pb.BackendClient, error) {

	opts := []grpc.DialOption{}
	switch secure {
	case true:
		creds, err := credentials.NewClientTLSFromFile(cert, "")
		if err != nil {
			return nil, errors.Wrapf(err, "could not load tls cert: %s", cert)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	case false:
		opts = append(opts, grpc.WithInsecure())
	}
	// opts = append(opts, grpc.WithInsecure())

	log.Debugf("dialing to %s\n", serverAddr)
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "new client")
	}

	return pb.NewBackendClient(conn), nil
}
