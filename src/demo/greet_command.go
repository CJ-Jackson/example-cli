package demo

import (
	"fmt"
	"github.com/cjtoolkit/cli"
	"github.com/cjtoolkit/cli/options"
	"strings"
)

type greetCommand struct {
	text string
	yell bool
}

func (gC *greetCommand) CommandConfigure(c *cli.Command) {
	c.SetName("demo:greet").
		SetDescription("Greet someone").
		AddOption("name", "Who do you want to greet?", options.NewString(&gC.text, options.StringMaxRune(50))).
		AddOption("yell", "If set, the task will yell in uppercase letters", options.NewBool(&gC.yell))
}

func (gC *greetCommand) CommandExecute() {
	if "" == gC.text {
		gC.text = "Hello!"
	} else {
		gC.text = "Hello " + gC.text + "!"
	}

	if gC.yell {
		gC.text = strings.ToUpper(gC.text)
	}

	fmt.Println(gC.text)
}

func init() {
	cli.RegisterCommand(&greetCommand{})
}
