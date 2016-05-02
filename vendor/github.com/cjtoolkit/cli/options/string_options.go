package options

import "regexp"

func StringMinRune(minRune int) func(*String) {
	return func(s *String) {
		s.MinRune = minRune
	}
}

func StringMaxRune(maxRune int) func(*String) {
	return func(s *String) {
		s.MaxRune = maxRune
	}
}

func StringPattern(pattern *regexp.Regexp) func(*String) {
	return func(s *String) {
		s.Pattern = pattern
	}
}