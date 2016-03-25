package help

/*
Implement:
	Interface in "sort"
*/
type Options []Option

func (o Options) Len() int {
	return len(o)
}

func (o Options) Less(i, j int) bool {
	return o[i].Name < o[j].Name
}

func (o Options) Swap(i, j int) {
	o[i], o[j] = o[j], o[i]
}

func (o Options) Plural() (plural string) {
	if len(o) > 1 {
		plural = "s"
	}
	return
}
