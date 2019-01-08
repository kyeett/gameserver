package game

import (
	"flag"
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
	// {"remote-state",
	// 	[]Option{RemoteState("localhost:10001", false), DevServer("localhost:10001", false)},
	// },
	{"remote-state-secure",
		[]Option{RemoteState("localhost:10002", true), DevServer("localhost:10002", true)},
	},
}

func Test_NewPlayers(t *testing.T) {
	nPlayers := 100 //

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

var stress = flag.Bool("stress", false, "run stress tests")

func Test_StressTest(t *testing.T) {
	if !*stress {
		t.Skip("Stress test skipped, use -stress flag")
	}
	log.SetLevel(log.DebugLevel)

	nGames := 5
	for _, secure := range []bool{false, true} {
		for i := 0; i < nGames; i++ {
			t.Run("remote-new-game-secure:"+strconv.FormatBool(secure), func(t *testing.T) {
				t.Parallel()
				p := freeTestPort(t)
				host := "localhost:" + p
				opts := []Option{RemoteState(host, secure), DevServer(host, secure)}
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
}

func freeTestPort(t *testing.T) string {
	port, err := freeport.GetFreePort()
	if err != nil {
		t.Fatal(err)
	}

	return strconv.Itoa(port)
}
