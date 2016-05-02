package cli

type CommandInterface interface {
	CommandConfigure(c *Command)
	CommandExecute()
}

type CommandPreInterface interface {
	CommandInterface
	CommandPre()
}

func execPre(c CommandInterface) {
	if c, ok := c.(CommandPreInterface); ok {
		c.CommandPre()
	}
}