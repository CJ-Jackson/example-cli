package help

import (
	"unicode/utf8"
)

func spacer(value string, maxSpace int) (space string) {
	space = "  "
	numOfSpaces := maxSpace - utf8.RuneCountInString(value)
	for i := 0; i < numOfSpaces; i++ {
		space += " "
	}
	return
}

func bigValueCheckAndUpdate(value string, countOfBiggestValue *int) {
	if count := utf8.RuneCountInString(value); count > *countOfBiggestValue {
		*countOfBiggestValue = count
	}
}
