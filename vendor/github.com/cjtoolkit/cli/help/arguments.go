package help

type Arguments []Argument

func (a Arguments) Plural() (plural string) {
	if len(a) > 1 {
		plural = "s"
	}
	return
}
