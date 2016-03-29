package global

import (
	"github.com/cjtoolkit/cli"
	"github.com/cjtoolkit/cli/options"
)

/*
Implement:
	GlobalInterface in "githun.com/cjtoolkit/cli"
*/
type globalGlobal struct{}

func (_ globalGlobal) GlobalConfigure(g *cli.Global) {
	g.AddOption("prod", "Set to Production Mode", options.Bool{Ptr: &global.Prod})
}

func (_ globalGlobal) Lock() {
	globalSync.Lock()
}

func (_ globalGlobal) Unlock() {
	globalSync.Unlock()
}

func init() {
	cli.RegisterGlobal(globalGlobal{})
}
