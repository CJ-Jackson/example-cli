package options

func IntMin(min int64) func(*Int) {
	return func(i *Int) {
		i.Min = min
		i.MinZero = true
	}
}

func IntMax(max int64) func(*Int) {
	return func(i *Int) {
		i.Max = max
		i.MaxZero = true
	}
}