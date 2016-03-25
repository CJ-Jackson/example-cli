package cli

import (
	"regexp"
)

var (
	commandNamePattern = regexp.MustCompile(`^[^\s:-]{1}[^\s:]+[:]{1}[^\s]+$`)
	otherNamePattern   = regexp.MustCompile(`^[^\s-=]{1}[^\s=]*$`)
)
