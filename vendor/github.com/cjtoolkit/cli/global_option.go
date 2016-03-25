package cli

type globalOption struct {
	Name        string
	Constaint   string
	Description string
	Transformer OptionTransformerInterface
}

func (gO *globalOption) postCheck() {
	switch {
	case "" == gO.Name:
		panic("'name' cannot be left blank")
	case !otherNamePattern.MatchString(gO.Name):
		panic("'name' does not match `" + otherNamePattern.String() + "`")
	case nil == gO.Transformer:
		panic(gO.Name + ": 'transformer' cannot be nil")
	default:
		defer handleErrorAndPanicAgain(gO.Name + ": Transformer: ")
		gO.Transformer.PreCheck()
		gO.Constaint = gO.Transformer.Constaint()
	}
}

func (gO *globalOption) populate(op *options) {
	defer handleErrorAndPanicAgain("Option: " + gO.Name + ": ")
	op.setName(gO.Name)
	gO.Transformer.OptionTransform(op)
}
