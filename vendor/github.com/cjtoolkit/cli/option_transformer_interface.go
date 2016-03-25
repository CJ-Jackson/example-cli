package cli

type OptionTransformerInterface interface {
	PreCheck()
	Constaint() string
	OptionTransform(option OptionsInterface)
}
