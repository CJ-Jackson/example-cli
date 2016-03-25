package help

import (
	"io"
	"sort"
	"text/template"
)

type CommandHelp struct {
	countOfBiggestName        int
	countOfBiggestDescription int
	Name                      string
	Options                   Options
	MandatoryArguments        Arguments
	OptionalArguments         Arguments
}

const COMMAND_HELP_TEMPLATE = `{{$top := . -}}
Command Help ('{{.Name}}'):

{{print .Name " "}}
{{- if .HasMandatoryArguments -}}
{{- range .MandatoryArguments}}{{.Name}} {{end -}}
{{- end}}
{{- if .HasOptionalArguments}}{{"[ " -}}
{{- range .OptionalArguments}}{{.Name}} {{end}}{{- "]" -}}
{{- end}}

{{- if .HasOptions}}

Option{{.Options.Plural}}: (They do not count as argument.)

  Name{{"Name"|.NameSpacer}}Description{{"Description"|.DescriptionSpacer}}Constraint
{{range .Options}}{{"  " -}}
--{{.Name}}{{print "--" .Name|$top.NameSpacer}}{{.Description}}{{.Description|$top.DescriptionSpacer -}}
{{.Constraint}} Mandatory:'{{.Mandatory}}'
{{end -}}
{{end -}}
{{- if .HasMandatoryArguments}}

Mandatory Argument{{.MandatoryArguments.Plural}}:

  Name{{"Name"|.NameSpacer}}Description{{"Description"|.DescriptionSpacer}}Constraint
{{range .MandatoryArguments}}{{"  " -}}
{{.Name}}{{.Name|$top.NameSpacer}}{{.Description}}{{.Description|$top.DescriptionSpacer -}}
{{.Constraint}}
{{end -}}
{{end -}}
{{- if .HasOptionalArguments}}

Optional Argument{{.OptionalArguments.Plural}}:

  Name{{"Name"|.NameSpacer}}Description{{"Description"|.DescriptionSpacer}}Constraint
{{range .OptionalArguments}}{{"  " -}}
{{.Name}}{{.Name|$top.NameSpacer}}{{.Description}}{{.Description|$top.DescriptionSpacer -}}
{{.Constraint}}
{{end -}}
{{end}}

`

func NewCommandHelp(commandName string) *CommandHelp {
	return &CommandHelp{
		Name:                      commandName,
		countOfBiggestName:        4,
		countOfBiggestDescription: 11,
	}
}

func (cH *CommandHelp) HasOptions() bool {
	return nil != cH.Options
}

func (cH *CommandHelp) HasMandatoryArguments() bool {
	return nil != cH.MandatoryArguments
}

func (cH *CommandHelp) HasOptionalArguments() bool {
	return nil != cH.OptionalArguments
}

func (cH *CommandHelp) NameSpacer(value string) string {
	return spacer(value, cH.countOfBiggestName)
}

func (cH *CommandHelp) DescriptionSpacer(value string) string {
	return spacer(value, cH.countOfBiggestDescription)
}

func (cH *CommandHelp) Finalise() {
	cH.finaliseOptions()
	cH.finaliseMandatoryArgument()
	cH.finaliseOptionalArgument()
}

func (cH *CommandHelp) finaliseOptions() bool {
	if !cH.HasOptions() {
		return false
	}

	sort.Sort(cH.Options)

	for _, op := range cH.Options {
		bigValueCheckAndUpdate("--"+op.Name, &cH.countOfBiggestName)
		bigValueCheckAndUpdate(op.Description, &cH.countOfBiggestDescription)
	}

	return true
}

func (cH *CommandHelp) finaliseMandatoryArgument() bool {
	if !cH.HasMandatoryArguments() {
		return false
	}

	cH.finaliseArgument(cH.MandatoryArguments)

	return true
}

func (cH *CommandHelp) finaliseOptionalArgument() bool {
	if !cH.HasOptionalArguments() {
		return false
	}

	cH.finaliseArgument(cH.OptionalArguments)

	return true
}

func (cH *CommandHelp) finaliseArgument(Arguments Arguments) {
	for _, arg := range Arguments {
		bigValueCheckAndUpdate(arg.Name, &cH.countOfBiggestName)
		bigValueCheckAndUpdate(arg.Description, &cH.countOfBiggestDescription)
	}
}

func (cH *CommandHelp) Render(w io.Writer) {
	template.Must(template.New("Command-Help").Parse(COMMAND_HELP_TEMPLATE)).
		Execute(w, cH)
}
