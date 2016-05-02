package arguments

import (
	"fmt"
	"regexp"
	"unicode/utf8"
)

/*
Implement:
	ArgumentTransformerInterface in "github.com/cjtoolkit/cli"
*/
type String struct {
	Ptr     *string // Mandatory
	MinRune int
	MaxRune int
	Pattern *regexp.Regexp
}

func NewString(ptr *string, options ...func(*String)) String {
	s := String{Ptr: ptr}

	for _, option := range options {
		option(&s)
	}

	return s
}

func (s String) PreCheck() {
	switch {
	case nil == s.Ptr:
		panic("Ptr cannot be nil")
	}
}

func (s String) Constraint() string {
	str := "Type:'string'"

	if s.MinRune > 0 {
		str += fmt.Sprint(" Min:'", s.MinRune, "'")
	}

	if s.MaxRune > 0 {
		str += fmt.Sprint(" Max:'", s.MaxRune, "'")
	}

	if nil != s.Pattern {
		str += " Pattern:'" + s.Pattern.String() + "'"
	}

	return str
}

func (s String) ArgumentTransform(argument string) {
	*s.Ptr = argument

	s.validate()
}

func (s String) validate() {
	s.validateMinRune()
	s.validateMaxRune()
	s.validatePattern()
}

func (s String) validateMinRune() {
	switch {
	case 0 == s.MinRune:
		return
	case s.MinRune > utf8.RuneCountInString(*s.Ptr):
		panic(fmt.Sprintf("Should be more than '%d' in length", s.MinRune))
	}
}

func (s String) validateMaxRune() {
	switch {
	case 0 == s.MaxRune:
		return
	case s.MaxRune < utf8.RuneCountInString(*s.Ptr):
		panic(fmt.Sprintf("Should be less than '%d' in length", s.MaxRune))
	}
}

func (s String) validatePattern() {
	switch {
	case nil == s.Pattern:
		return
	case !s.Pattern.MatchString(*s.Ptr):
		panic(fmt.Sprintf("Should match '%s'", s.Pattern.String()))
	}
}
