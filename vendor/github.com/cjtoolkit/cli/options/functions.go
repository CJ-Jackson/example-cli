package options

func CreatePanicFunctionOnFalse(value bool, name string) func() {
	if !value {
		return func() {
			panic("'" + name + "' is mandatory.")
		}
	}
	return nil
}
