package cli

import (
	"github.com/cjtoolkit/cli/help"
	"strings"
)

type Command struct {
	command           CommandInterface
	name              string
	description       string
	options           []*commandOption
	mandatoryArgument []*commandArgument
	optionalArgument  []*commandArgument
}

func newCommand(command CommandInterface) *Command {
	return &Command{
		command:           command,
		options:           []*commandOption{},
		mandatoryArgument: []*commandArgument{},
		optionalArgument:  []*commandArgument{},
	}
}

func (c *Command) SetName(name string) *Command {
	c.name = name
	return c
}

func (c *Command) SetDescription(description string) *Command {
	c.description = description
	return c
}

func (c *Command) AddOption(
	name, description string,
	mandatory bool,
	transformer OptionTransformerInterface,
) *Command {
	c.options = append(c.options, &commandOption{
		Name:        strings.TrimSpace(name),
		Description: strings.TrimSpace(description),
		Mandatory:   mandatory,
		Transformer: transformer,
	})
	return c
}

func (c *Command) AddArgument(
	name, description string,
	mandatory bool,
	transformer ArgumentTransformerInterface,
) *Command {
	argument := &commandArgument{
		Name:        strings.TrimSpace(name),
		Description: strings.TrimSpace(description),
		Mandatory:   mandatory,
		Transformer: transformer,
	}
	execTrueFalse(mandatory, func() {
		c.mandatoryArgument = append(c.mandatoryArgument, argument)
	}, func() {
		c.optionalArgument = append(c.optionalArgument, argument)
	})
	return c
}

func (c *Command) postCheck() {
	switch {
	case "" == c.name:
		panic("Name cannot be left blank")
	case !commandNamePattern.MatchString(c.name):
		panic("Name does not match `" + commandNamePattern.String() + "`")
	default:
		c.checkAllOptions()
		c.checkAllArgument()
	}
}

func (c *Command) checkAllOptions() {
	defer handleErrorAndPanicAgain("Command Options: " + c.name + ": ")
	for _, option := range c.options {
		option.postCheck()
	}
}

func (c *Command) checkAllArgument() {
	defer handleErrorAndPanicAgain("Command Argument: " + c.name + ": ")
	for _, argument := range append(c.mandatoryArgument, c.optionalArgument...) {
		argument.postCheck()
	}
}

func (c *Command) execCommand(argOp argumentOptions) {
	c.populateOptions(argOp.options)
	argOp.options.checkForUnrecognisedOption()
	c.populateArgument(argOp.arguments[1:])
	c.command.CommandExecute()
}

func (c *Command) populateOptions(op *options) {
	for _, option := range c.options {
		option.populate(op)
	}
}

func (c *Command) populateArgument(arguments []string) {
	argumentsCount := len(arguments)
	for key, argument := range append(c.mandatoryArgument, c.optionalArgument...) {
		if key >= argumentsCount {
			if argument.Mandatory {
				panic("Argument: " + argument.Name + " is required.")
			} else {
				break
			}
		}
		argument.populate(arguments[key])
	}
}

func (c *Command) collectGeneralHelp(helpData *help.CommandHelp) {
	for _, op := range c.options {
		helpData.Options = append(helpData.Options, help.Option{
			Name:        op.Name,
			Description: op.Description,
			Constraint:  op.Constaint,
			Mandatory:   op.Mandatory,
		})
	}
	for _, arg := range c.mandatoryArgument {
		helpData.MandatoryArguments = append(helpData.MandatoryArguments, help.Argument{
			Name:        arg.Name,
			Description: arg.Description,
			Constraint:  arg.Constraint,
		})
	}
	for _, arg := range c.optionalArgument {
		helpData.OptionalArguments = append(helpData.OptionalArguments, help.Argument{
			Name:        arg.Name,
			Description: arg.Description,
			Constraint:  arg.Constraint,
		})
	}
}
