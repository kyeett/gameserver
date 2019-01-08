package game

import (
	"context"
	"strings"
	"time"

	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"

	"github.com/kyeett/gameserver"
	"github.com/kyeett/gameserver/grpc"
	"github.com/kyeett/gameserver/localstate"
	"github.com/kyeett/gameserver/types"
	"github.com/kyeett/mapgenerator/gen"
)

type Game struct {
	ctx  context.Context
	opts options
	gameserver.GameState
}

// An option sets options such as what type of world to use, number of players, etc.
type Option func(*options) error

type options struct {
	world             types.World
	initiateStateFunc func(*Game) error
	startDevServer    func(*Game) error
}

func initiateLocalState(g *Game) error {
	log.Debugf("Syncing locally")
	g.GameState = localstate.New(g.opts.world)
	return nil
}

var defaultGameOptions = options{
	world:             types.FirstWorld,
	initiateStateFunc: initiateLocalState,
	startDevServer:    func(_ *Game) error { return nil },
}

func New(opt ...Option) (*Game, error) {
	log.SetLevel(log.DebugLevel)
	filenameHook := filename.NewHook()
	filenameHook.Field = "source"
	log.AddHook(filenameHook)

	g := &Game{
		ctx:  context.Background(),
		opts: defaultGameOptions,
	}
	for _, o := range opt {
		err := o(&g.opts)
		if err != nil {
			return nil, err
		}
	}

	err := g.opts.startDevServer(g)
	if err != nil {
		return nil, err
	}

	err = g.opts.initiateStateFunc(g)
	if err != nil {
		return nil, err
	}

	log.Infof("Game created, starting\n")
	return g, nil
}

func World(name string) Option {

	return func(o *options) error {

		switch name {
		case "generate", "generate10,10":
			o.world = generateMap(10, 10)
		case "generate20x20":
			o.world = generateMap(20, 20)
		case "generate40x30":
			o.world = generateMap(40, 30)
		case "first":
			o.world = types.FirstWorld
		case "test4x4":
			o.world = types.Test4x4World
		case "big":
			o.world = types.BigWorld
		default:

			s := []string{"first", "test4x4", "big", "generate10,10", "generate20x20", "generate40,30"}
			return errors.Errorf("No such map '%s'. Valid ones are %s", name, strings.Join(s, ", "))
		}

		return nil
	}
}

func RemoteState(addr string, secure bool) Option {
	return func(o *options) error {
		log.Debugf("Configure state server to: remote\n")
		o.initiateStateFunc = func(g *Game) error {
			// serverAddr := fmt.Sprintf("localhost:%d", defaultPort)
			serverAddr := addr
			s, err := grpc.NewClient(serverAddr, secure)
			if err != nil {
				return err
			}
			g.GameState = s
			log.Debugf("Syncing state towards %s\n", serverAddr)
			return nil
		}
		return nil
	}
}

// DevServer creates and runs game state server on localhost
func DevServer(host string) Option {
	return func(o *options) error {
		log.Debugf("Configure game to start local dev server\n")
		o.startDevServer = func(g *Game) error {
			log.Debugf("Start dev server in background\n")

			//Todo: don't start go routine in here?
			ss, err := grpc.NewServer(g.opts.world)
			if err != nil {
				return err
			}

			go ss.Run(g.ctx, host)

			return nil
		}
		return nil
	}
}

func generateMap(width, height int) types.World {
	m := gen.GenerateParam(2.5, 2.5, time.Now().Unix(), 2, width, height)

	tiles := []types.Tile{}

	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			if m[y][x] < 0.4 {
				tiles = append(tiles, types.Water)
			} else {
				tiles = append(tiles, types.Grass)
			}
		}
	}
	return types.NewWorld(tiles, width, height)
}
