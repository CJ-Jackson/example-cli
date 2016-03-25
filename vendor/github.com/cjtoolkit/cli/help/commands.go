package help

type Commands map[string]*Command

func (c Commands) GetCommand(name string) *Command {
	return c[name]
}

func (c Commands) Plural() (plural string) {
	if len(c) > 1 {
		plural = "s"
	}
	return
}
