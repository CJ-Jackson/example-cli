package help

type CommandNames map[string][]string

func (cN CommandNames) GetCommandNames(name string) []string {
	return cN[name]
}
