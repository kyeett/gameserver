package game

import (
	"strconv"
	"testing"

	"github.com/kyeett/gameserver/types"
	"github.com/phayes/freeport"
	log "github.com/sirupsen/logrus"
)

var stateTypes = []struct {
	name string
	opts []Option
}{
	{"local",
		[]Option{}},
	{"remote-state",
		[]Option{RemoteState("localhost:10001"), DevServer("10001")}},
}

func Test_NewPlayers(t *testing.T) {
	nPlayers := 1 //

	log.SetLevel(log.DebugLevel)
	for _, tc := range stateTypes {

		tcOpts := tc.opts
		t.Run(tc.name+"-new-players", func(t *testing.T) {
			t.Parallel()
			g, err := New(tcOpts...)
			if err != nil {
				t.Fatalf("creating game failed: %s\n", err)
			}

			// Save number of entitites before operations
			// Make sure server is up
			g.NewPlayer()

			nBefore := len(g.Entities())
			// Connect multiple new players
			for i := 0; i < nPlayers; i++ {
				g.NewPlayer()
			}

			if len(g.Entities()) != nBefore+nPlayers {
				t.Fatalf("expected %d entities, got %d", nBefore+nPlayers, len(g.Entities()))
			}
		})
	}
}

func Test_StressTest(t *testing.T) {

	log.SetLevel(log.DebugLevel)

	nGames := 100
	for i := 0; i < nGames; i++ {
		t.Run("remote-new-game", func(t *testing.T) {
			t.Parallel()
			p := freeTestPort(t)
			opts := []Option{RemoteState("localhost:" + p), DevServer(p)}
			g, err := New(opts...)
			if err != nil {
				t.Fatalf("creating game failed: %s\n", err)
			}

			e, err := g.NewPlayer()
			if err != nil {
				t.Fatal(p)
			}

			for m := 0; m < 100; m++ {
				pos := types.Position{Coord: types.Coord{X: m%2 + 1, Y: m%2 + 1}, Theta: 3}
				_, err := g.PerformAction(e, pos)
				if err != nil {
					t.Fatal(err)
				}
			}
		})
	}
}

func freeTestPort(t *testing.T) string {
	port, err := freeport.GetFreePort()
	if err != nil {
		t.Fatal(err)
	}

	return strconv.Itoa(port)
}
