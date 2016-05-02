package arguments

func FloatMin(min float64) func(*Float) {
	return func(f *Float) {
		f.Min = min
		f.MinZero = true
	}
}

func FloatMax(max float64) func(*Float) {
	return func(f *Float) {
		f.Max = max
		f.MaxZero = true
	}
}