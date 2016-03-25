package cli

type OptionsInterface interface {
	GetName() string
	ExecOnMandatory(fn func())
	HasOne() bool
	GetOne() string
	GetAll() []string
}
