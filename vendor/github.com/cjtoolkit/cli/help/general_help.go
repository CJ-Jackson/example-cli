package help

import (
	"io"
	"sort"
	"strings"
	"text/template"
)

type GeneralHelp struct {
	countOfBiggestName        int
	countOfBiggestDescription int
	GlobalsOptions            Options
	Commands                  Commands
	TopCommand                []string
	CommandNames              CommandNames
}

const GENERAL_HELP_TEMPLATE = `{{$top := . -}}
General Help:

{{- if .HasGlobalOptions}}

Global Option{{.GlobalsOptions.Plural}}: (They do not count as argument.)

  Name{{"Name"|.NameSpacer}}Description{{"Description"|.DescriptionSpacer}}Constraint
{{range .GlobalsOptions}}{{"  " -}}
--{{.Name}}{{print "--" .Name|$top.NameSpacer}}{{.Description}}{{.Description|$top.DescriptionSpacer -}}
{{.Constraint}}
{{end -}}
{{else}}
{{end}}
{{if .HasCommands -}}
Command{{.Commands.Plural}}: (Add '--help' before or after the command name, to get more details.)

  Name{{"Name"|.NameSpacer}}Description
{{range .TopCommand -}}
{{.}}:
{{range $top.CommandNames.GetCommandNames . -}}
{{$command := $top.Commands.GetCommand . }}{{"  " -}}
{{$command.Name}}{{$command.Name|$top.NameSpacer}}{{$command.Description}}
{{end -}}
{{end -}}
{{end}}

`

func NewGeneralHelp() *GeneralHelp {
	return &GeneralHelp{
		countOfBiggestName:        4,
		countOfBiggestDescription: 11,
		TopCommand:                []string{},
		CommandNames:              CommandNames{},
	}
}

func (gH *GeneralHelp) HasGlobalOptions() bool {
	return nil != gH.GlobalsOptions
}

func (gH *GeneralHelp) HasCommands() bool {
	return nil != gH.Commands
}

func (gH *GeneralHelp) NameSpacer(value string) string {
	return spacer(value, gH.countOfBiggestName)
}

func (gH *GeneralHelp) DescriptionSpacer(value string) (space string) {
	return spacer(value, gH.countOfBiggestDescription)
}

func (gH *GeneralHelp) Finalise() {
	gH.finaliseGlobalOptions()
	gH.finaliseCommands()
}

func (gH *GeneralHelp) finaliseGlobalOptions() bool {
	if !gH.HasGlobalOptions() {
		return false
	}

	sort.Sort(gH.GlobalsOptions)

	for _, op := range gH.GlobalsOptions {
		bigValueCheckAndUpdate("--"+op.Name, &gH.countOfBiggestName)
		bigValueCheckAndUpdate(op.Description, &gH.countOfBiggestDescription)
	}

	return true
}

func (gH *GeneralHelp) finaliseCommands() bool {
	if !gH.HasCommands() {
		return false
	}

	for name, _ := range gH.Commands {
		bigValueCheckAndUpdate(name, &gH.countOfBiggestName)
		top := name[:strings.Index(name, ":")]
		gH.CommandNames[top] = append(gH.CommandNames[top], name)
	}

	for top, sortValue := range gH.CommandNames {
		gH.TopCommand = append(gH.TopCommand, top)
		sort.Sort(sort.StringSlice(sortValue))
		gH.CommandNames[top] = sortValue
	}

	sort.Sort(sort.StringSlice(gH.TopCommand))

	return true
}

func (gH *GeneralHelp) Render(w io.Writer) {
	template.Must(template.New("General-Help").Parse(GENERAL_HELP_TEMPLATE)).
		Execute(w, gH)
}
