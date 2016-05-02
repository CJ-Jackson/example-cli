package options

import (
	"github.com/cjtoolkit/cli"
)

/*
Implement:
	OptionTransformerInterface in "github.com/cjtoolkit/cli"
*/
type Bool struct {
	Ptr *bool // Mandatory
}

func NewBool(ptr *bool) Bool {
	return Bool{Ptr: ptr}
}

func (b Bool) PreCheck() {
	if nil == b.Ptr {
		panic("Ptr cannot be nil")
	}
}

func (b Bool) Constraint() string {
	return "Type:'bool'"
}

func (b Bool) OptionTransform(option cli.OptionsInterface) {
	*b.Ptr = !*b.Ptr
}
