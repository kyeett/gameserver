package grpc

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/kyeett/gameserver/types"
)

var (
	defaultPort = "10003"
	defaultAddr = fmt.Sprintf("localhost:%s", defaultPort)
	secure      = true
)

func Test_NewPlayers(t *testing.T) {
	nPlayers := 100

	ctx := context.Background()
	ctx, cancelServer := context.WithCancel(ctx)
	s, err := NewServer(types.Test4x4World)
	if err != nil {
		t.Fatal("Server did not start properly", err)
	}
	fmt.Println(s.local.World())
	time.Sleep(1 * time.Millisecond)
	go s.Run(ctx, defaultPort, secure)
	fmt.Println(s.local.World())
	time.Sleep(1 * time.Millisecond)

	c, err := NewClient(defaultAddr, secure)
	if err != nil {
		t.Fatal("new client", err)
	}

	// Save number of entitites before operations
	nBefore := len(c.Entities())

	// Connect multiple new players
	wg := sync.WaitGroup{}
	for i := 0; i < nPlayers; i++ {
		wg.Add(1)
		go func() {

			defer wg.Done()
			testClient, err := NewClient(defaultAddr, secure)
			if err != nil {
				t.Fatal("new client", err)
			}

			_, err = testClient.NewPlayer()
			if err != nil {
				t.Fatal("new player", err)
			}

			// Keep clients connected
			time.Sleep(100 * time.Millisecond)
		}()
	}
	wg.Wait()

	if len(c.Entities()) != nBefore+nPlayers {
		t.Fatalf("expected %d entities, got %d", nBefore+nPlayers, len(c.Entities()))
	}

	cancelServer()
}
