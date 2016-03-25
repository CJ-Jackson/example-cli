package cli

type ArgumentTransformerInterface interface {
	PreCheck()
	Constaint() string
	ArgumentTransform(argument string)
}
