package date

import (
	"fmt"
	"github.com/cjtoolkit/cli"
	"github.com/cjtoolkit/cli/arguments"
	"regexp"
	"time"
)

type ageCommand struct {
	name string
	dob  time.Time
	now  time.Time
}

func (aC *ageCommand) CommandConfigure(c *cli.Command) {
	aC.now = time.Now()

	c.SetName("date:dob").
		SetDescription("Calculate Date of Birth").
		AddArgument("name", "Your Name", arguments.NewString(&aC.name,
			arguments.StringMinRune(5), arguments.StringPattern(regexp.MustCompile(`^[A-Za-z]*$`)))).
		AddArgument("dob", "Your date of birth", arguments.NewTime(&aC.dob, "02-01-2006",
			arguments.TimeMax(aC.now)))
}

func (aC *ageCommand) CommandExecute() {
	duration := aC.now.Sub(aC.dob)
	fmt.Printf("Hello, my name is '%s' and I'm '%d' years old", aC.name,
		time.Date(0, 0, 0, 0, 0, int(duration.Seconds()), 0, aC.now.Location()).Year())
	fmt.Println()
}

func init() {
	cli.RegisterCommand(&ageCommand{})
}
