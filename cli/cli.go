package cli // import "bazil.org/bazil/cli"

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/pprof"
	"sync"
	"time"

	"bazil.org/bazil/cliutil/flagx"
	"bazil.org/bazil/cliutil/subcommands"
	"bazil.org/bazil/defaults"
	"bazil.org/bazil/server/control/wire"
	"bazil.org/bazil/util/grpcunix"
	"bazil.org/fuse"
	"github.com/tv42/jog"
	"google.golang.org/grpc"
)

type bazil struct {
	flag.FlagSet
	Config struct {
		Verbose    bool
		Debug      bool
		DataDir    flagx.AbsPath
		CPUProfile string
	}
	Log *jog.Logger

	control struct {
		setup  sync.Once
		err    error
		conn   *grpc.ClientConn
		client wire.ControlClient
	}
}

var _ Service = (*bazil)(nil)

func (b *bazil) Setup() (ok bool) {
	b.Log = jog.New(nil)
	if b.Config.Debug {
		// regular FUSE debug happens through VolumeRef.debug, but (in
		// theory) this might see some events not related to an
		// bazil.org/fuse/fs Serve loop, so keep it around
		fuse.Debug = b.Log.Event
	}

	if b.Config.CPUProfile != "" {
		f, err := os.Create(b.Config.CPUProfile)
		if err != nil {
			log.Printf("cpu profiling: %v", err)
			return false
		}
		err = pprof.StartCPUProfile(f)
		if err != nil {
			log.Printf("cpu profiling: %v", err)
			return false
		}
	}
	return true
}

func (b *bazil) Teardown() (ok bool) {
	if b.Config.CPUProfile != "" {
		pprof.StopCPUProfile()
	}
	if b.control.conn != nil {
		b.control.conn.Close()
	}
	return true
}

// Control returns a gRPC client that can make control requests to the
// server. This allows dialing only when needed, because not every cli
// tool is a control client.
//
// TODO behavior if server is not there; grpc causes infinite retries
func (b *bazil) Control() (wire.ControlClient, error) {
	b.control.setup.Do(b.controlDial)
	return b.control.client, b.control.err
}

func (b *bazil) controlDial() {
	b.control.conn, b.control.err = grpcunix.Dial(
		filepath.Join(b.Config.DataDir.String(), "control"),
		grpc.WithTimeout(500*time.Millisecond),
	)
	b.control.client = wire.NewControlClient(b.control.conn)
}

// Bazil allows command-line callables access to global flags, such as
// verbosity.
var Bazil = bazil{}

func init() {
	Bazil.BoolVar(&Bazil.Config.Verbose, "v", false, "verbose output")
	Bazil.BoolVar(&Bazil.Config.Debug, "debug", false, "debug output")

	Bazil.Config.DataDir = flagx.AbsPath(defaults.DataDir())
	// ensure absolute path to make the control socket show up nicer
	// in `ss` output
	Bazil.Var(&Bazil.Config.DataDir, "data-dir", "path to filesystem state")

	Bazil.StringVar(&Bazil.Config.CPUProfile, "cpuprofile", "", "write cpu profile to file")

	subcommands.Register(&Bazil)
}

// Service is an interface that commands can implement to setup and
// teardown services for the subcommands below them.
//
// As Run and potential multiple Teardown failures makes having a
// single error return impossible, Setup and Teardown only get to
// signal a boolean success. Any detail should be exposed via log.
type Service interface {
	Setup() (ok bool)
	Teardown() (ok bool)
}

func run(result subcommands.Result) (ok bool) {
	var cmd interface{}
	for _, cmd = range result.ListCommands() {
		if svc, isService := cmd.(Service); isService {
			ok = svc.Setup()
			if !ok {
				return false
			}
			defer func() {
				// Teardown failures can cause non-successful exit
				if !svc.Teardown() {
					ok = false
				}
			}()
		}
	}
	run := cmd.(subcommands.Runner)
	err := run.Run()
	if err != nil {
		log.Printf("error: %v", err)
		return false
	}
	return true
}

// Main is primary entry point into the bazil command line
// application.
func Main() (exitstatus int) {
	progName := filepath.Base(os.Args[0])
	log.SetFlags(0)
	log.SetPrefix(progName + ": ")

	result, err := subcommands.Parse(&Bazil, progName, os.Args[1:])
	if err == flag.ErrHelp {
		result.Usage()
		return 0
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", result.Name(), err)
		result.Usage()
		return 2
	}

	ok := run(result)
	if !ok {
		return 1
	}
	return 0
}
