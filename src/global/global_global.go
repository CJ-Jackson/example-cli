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
	g.AddOption("prod", "Set to Production Mode", options.NewBool(&global.Prod))
}

func init() {
	cli.RegisterGlobal(globalGlobal{})
}
