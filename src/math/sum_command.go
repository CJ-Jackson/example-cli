package math

import (
	"fmt"
	"github.com/cjtoolkit/cli"
	"github.com/cjtoolkit/cli/arguments"
)

type sumCommand struct {
	number1 int64
	number2 int64
}

func (sC *sumCommand) CommandConfigure(c *cli.Command) {
	c.SetName("math:sum").
		SetDescription("Add two numbers together").
		AddArgument("Number1", "", arguments.Int{Ptr: &sC.number1, MinZero: true}).
		AddArgument("Number2", "", arguments.Int{Ptr: &sC.number2, MinZero: true})
}

func (sC *sumCommand) CommandExecute() {
	fmt.Println(sC.number1 + sC.number2)
}

func init() {
	cli.RegisterCommand(&sumCommand{})
}
