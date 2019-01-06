package grpc

import (
	"context"
	"sync"
	"testing"
	"time"
)

func Test_NewPlayers(t *testing.T) {
	nPlayers := 100

	ctx := context.Background()
	ctx, cancelServer := context.WithCancel(ctx)
	s := NewServer()
	go s.Run(ctx)

	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	// Save number of entitites before operations
	nBefore := len(c.Entities())

	// Connect multiple new players
	wg := sync.WaitGroup{}
	for i := 0; i < nPlayers; i++ {
		wg.Add(1)
		go func() {

			defer wg.Done()
			testClient, err := NewClient()
			if err != nil {
				t.Fatal(err)
			}

			_, err = testClient.NewPlayer()
			if err != nil {
				t.Fatal(err)
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

func Test_Move(t *testing.T) {
	ctx := context.Background()
	ctx, cancelServer := context.WithCancel(ctx)
	s := NewServer()
	go s.Run(ctx)

	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	ec.NewPlayer()

	// Save number of entitites before operations
	nBefore := len(c.Entities())

	// Connect multiple new players
	wg := sync.WaitGroup{}
	for i := 0; i < nPlayers; i++ {
		wg.Add(1)
		go func() {

			defer wg.Done()
			testClient, err := NewClient()
			if err != nil {
				t.Fatal(err)
			}

			_, err = testClient.NewPlayer()
			if err != nil {
				t.Fatal(err)
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
