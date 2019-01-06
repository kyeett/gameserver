package main

import (
	"context"
	"log"

	"github.com/kyeett/gameserver/types"

	"github.com/kyeett/gameserver/grpc"
)

func main() {

	ss, err := grpc.NewServer(types.FirstWorld)
	if err != nil {
		log.Fatal(err)
	}
	ss.Run(context.Background())

}
