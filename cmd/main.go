package main

import (
	"context"
	"flag"

	log "github.com/sirupsen/logrus"

	"github.com/kyeett/gameserver/types"
	"github.com/onrik/logrus/filename"

	"github.com/kyeett/gameserver/grpc"
)

func main() {
	var host string
	var secure bool
	flag.StringVar(&host, "host", "localhost:10001", "hostname including port")
	flag.BoolVar(&secure, "secure", false, "enable TLS")
	flag.Parse()

	//Enable tracing
	log.SetLevel(log.DebugLevel)
	filenameHook := filename.NewHook()
	filenameHook.Field = "source"
	log.AddHook(filenameHook)

	ss, err := grpc.NewServer(types.FirstWorld)
	if err != nil {
		log.Fatal(err)
	}

	ss.Run(context.Background(), host, secure)
}
