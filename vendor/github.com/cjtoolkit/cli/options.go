package cli

import (
	"sort"
)

/*
Implement:
	OptionsInterface
*/
type options struct {
	options   map[string][]string
	count     map[string]int
	check     map[string]bool
	name      string
	mandatory bool
}

func newOption(op map[string][]string) *options {
	return &options{
		options: op,
		count:   map[string]int{},
		check:   map[string]bool{},
	}
}

func (o *options) setName(name string) {
	o.name = name
}

func (o *options) GetName() string {
	return o.name
}

func (o *options) setMandatory(mandatory bool) {
	o.mandatory = mandatory
}

func (o *options) ExecOnMandatory(fn func()) {
	if o.mandatory && nil != fn {
		fn()
	}
}

func (o *options) HasOne() bool {
	o.markCheck()
	return nil != o.options[o.name]
}

func (o *options) increment() {
	o.count[o.name]++
	o.markCheck()
}

func (o *options) GetOne() string {
	if nil == o.options[o.name] || o.count[o.name] >= len(o.options[o.name]) {
		return ""
	}
	defer o.increment()
	return o.options[o.name][o.count[o.name]]
}

func (o *options) delete() {
	o.count[o.name] = len(o.options[o.name])
	o.markCheck()
}

func (o *options) GetAll() []string {
	if nil == o.options[o.name] || o.count[o.name] >= len(o.options[o.name]) {
		return nil
	}
	defer o.delete()
	return o.options[o.name][o.count[o.name]:]
}

func (o *options) markCheck() {
	o.check[o.name] = true
}

func (o *options) checkForUnrecognisedOption() {
	var unrecognised []string

	for optionName, _ := range o.options {
		if !o.check[optionName] {
			unrecognised = append(unrecognised, optionName)
		}
	}

	if nil == unrecognised {
		return
	}

	sort.Sort(sort.StringSlice(unrecognised))

	msg := "Global or current command did not recognised those options:" + NEW_LINE + NEW_LINE

	for _, optionName := range unrecognised {
		msg += "--" + optionName + NEW_LINE
	}

	panic(msg)
}
