package game

import (
	"context"
	"strings"

	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"

	"github.com/kyeett/gameserver"
	"github.com/kyeett/gameserver/grpc"
	"github.com/kyeett/gameserver/localserver"
	"github.com/kyeett/gameserver/types"
)

type Game struct {
	ctx  context.Context
	opts options
	gameserver.GameServer
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
	g.GameServer = localserver.New(g.opts.world)
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
	opts := defaultGameOptions
	for _, o := range opt {
		err := o(&opts)
		if err != nil {
			return nil, err
		}
	}

	err := opts.startDevServer(g)
	if err != nil {
		return nil, err
	}

	err = opts.initiateStateFunc(g)
	if err != nil {
		return nil, err
	}

	log.Infof("Game created, starting\n")
	return g, nil
}

func World(name string) Option {
	definedWorlds := map[string]types.World{
		"first":   types.FirstWorld,
		"test4x4": types.Test4x4World,
	}

	return func(o *options) error {

		if _, ok := definedWorlds[name]; !ok {
			s := []string{}
			for world := range definedWorlds {
				s = append(s, world)
			}
			return errors.Errorf("No such map '%s'. Valid ones are %s", name, strings.Join(s, ", "))
		}

		o.world = definedWorlds[name]
		return nil
	}
}

func RemoteState(addr string, secure bool) Option {
	return func(o *options) error {
		log.Debugf("Configure state server to: remote\n")
		o.initiateStateFunc = func(g *Game) error {
			serverAddr := addr
			s, err := grpc.NewClient(serverAddr, secure)
			if err != nil {
				return err
			}
			g.GameServer = s
			log.Debugf("Syncing state towards %s\n", serverAddr)
			return nil
		}
		return nil
	}
}

// DevServer creates and runs game state server on localhost
func DevServer(port string, secure bool) Option {
	return func(o *options) error {
		log.Debugf("Configure game to start local dev server\n")
		o.startDevServer = func(g *Game) error {
			log.Debugf("Start dev server in background\n")

			//Todo: don't start go routine in here?
			ss, err := grpc.NewServer(g.opts.world)
			if err != nil {
				return err
			}

			go ss.Run(g.ctx, port, secure)

			return nil
		}
		return nil
	}
}
