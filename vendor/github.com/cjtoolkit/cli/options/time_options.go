package options

import "time"

func TimeMin(min time.Time) func(*Time) {
	return func(t *Time) {
		t.Min = min
		t.MinZero = true
	}
}

func TimeMax(max time.Time) func(*Time) {
	return func(t *Time) {
		t.Max = max
		t.MaxZero = true
	}
}