package math

import (
	"fmt"
	"github.com/cjtoolkit/cli"
	"github.com/cjtoolkit/cli/arguments"
)

type sumCli struct {
	number1 int64
	number2 int64
}

func (sC *sumCli) CommandConfigure(c *cli.Command) {
	c.SetName("math:sum").
		SetDescription("Add two numbers together").
		AddArgument("Number1", "", true, arguments.Int{Ptr: &sC.number1, MinZero: true}).
		AddArgument("Number2", "", true, arguments.Int{Ptr: &sC.number2, MinZero: true})
}

func (sC *sumCli) CommandExecute() {
	fmt.Println(sC.number1 + sC.number2)
}

func init() {
	cli.RegisterCommand(&sumCli{})
}
