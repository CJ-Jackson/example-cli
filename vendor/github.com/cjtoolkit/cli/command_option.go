package cli

type commandOption struct {
	Name        string
	Constaint   string
	Description string
	Mandatory   bool
	Transformer OptionTransformerInterface
}

func (cO *commandOption) postCheck() {
	switch {
	case "" == cO.Name:
		panic("'name' cannot be left blank")
	case !otherNamePattern.MatchString(cO.Name):
		panic("'name' does not match `" + otherNamePattern.String() + "`")
	case nil == cO.Transformer:
		panic(cO.Name + ": 'transformer' cannot be nil")
	default:
		defer handleErrorAndPanicAgain(cO.Name + ": Transformer: ")
		cO.Transformer.PreCheck()
		cO.Constaint = cO.Transformer.Constaint()
	}
}

func (cO *commandOption) populate(op *options) {
	defer handleErrorAndPanicAgain("Option: " + cO.Name + ": ")
	op.setName(cO.Name)
	op.setMandatory(cO.Mandatory)
	cO.Transformer.OptionTransform(op)
}
