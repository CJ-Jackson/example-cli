package demo

import (
	"fmt"
	"github.com/cjtoolkit/cli"
	"github.com/cjtoolkit/cli/options"
	"strings"
)

type greetCommand struct {
	name string
	yell bool
}

func (gC *greetCommand) CommandConfigure(c *cli.Command) {
	c.SetName("demo:greet").
		SetDescription("Greet someone").
		AddOption("name", "Who do you want to greet?",
			options.String{Ptr: &gC.name}).
		AddOption("yell", "If set, the task will yell in uppercase letters",
			options.Bool{Ptr: &gC.yell})
}

func (gC *greetCommand) CommandExecute() {
	if "" == gC.name {
		gC.name = "Hello!"
	} else {
		gC.name = "Hello " + gC.name + "!"
	}

	if gC.yell {
		gC.name = strings.ToUpper(gC.name)
	}

	fmt.Println(gC.name)
}

func init() {
	cli.RegisterCommand(&greetCommand{})
}
