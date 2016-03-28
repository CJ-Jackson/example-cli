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
	Arguments                 Arguments
}

const COMMAND_HELP_TEMPLATE = `{{$top := . -}}
Command Help ('{{.Name}}'):

Usage:
  {{print .Name " "}}
{{- if .HasArguments -}}
{{- range .Arguments}}{{.Name}} {{end -}}
{{- end}}

{{- if .HasOptions}}

Option{{.Options.Plural}}: (They do not count as argument)

  Name{{"Name"|.NameSpacer}}Description{{"Description"|.DescriptionSpacer}}Constraint
{{range .Options}}{{"  " -}}
--{{.Name}}{{print "--" .Name|$top.NameSpacer}}{{.Description}}{{.Description|$top.DescriptionSpacer -}}
{{.Constraint}}
{{end -}}
{{end -}}
{{- if .HasArguments}}

Argument{{.Arguments.Plural}}: (All of them are mandatory)

  Name{{"Name"|.NameSpacer}}Description{{"Description"|.DescriptionSpacer}}Constraint
{{range .Arguments}}{{"  " -}}
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

func (cH *CommandHelp) HasArguments() bool {
	return nil != cH.Arguments
}

func (cH *CommandHelp) NameSpacer(value string) string {
	return spacer(value, cH.countOfBiggestName)
}

func (cH *CommandHelp) DescriptionSpacer(value string) string {
	return spacer(value, cH.countOfBiggestDescription)
}

func (cH *CommandHelp) Finalise() {
	cH.finaliseOptions()
	cH.finaliseArgument()
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

func (cH *CommandHelp) finaliseArgument() bool {
	if !cH.HasArguments() {
		return false
	}

	for _, arg := range cH.Arguments {
		bigValueCheckAndUpdate(arg.Name, &cH.countOfBiggestName)
		bigValueCheckAndUpdate(arg.Description, &cH.countOfBiggestDescription)
	}

	return true
}

func (cH *CommandHelp) Render(w io.Writer) {
	template.Must(template.New("Command-Help").Parse(COMMAND_HELP_TEMPLATE)).
		Execute(w, cH)
}
