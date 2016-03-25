package global

import (
	"github.com/cjtoolkit/cli"
	"github.com/cjtoolkit/cli/options"
)

/*
Implement:
	GlobalInterface in "githun.com/cjtoolkit/cli"
*/
type globalCli struct{}

func (_ globalCli) GlobalConfigure(g *cli.Global) {
	g.AddOption("prod", "Set to Production Mode", options.Bool{Ptr: &global.Prod})
}

func (_ globalCli) Lock() {
	globalSync.Lock()
}

func (_ globalCli) Unlock() {
	globalSync.Unlock()
}

func init() {
	cli.RegisterGlobal(globalCli{})
}
