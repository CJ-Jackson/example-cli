package date

import (
	"fmt"
	"github.com/cjtoolkit/cli"
	"github.com/cjtoolkit/cli/arguments"
	"regexp"
	"time"
)

type ageCli struct {
	name string
	dob  time.Time
	now  time.Time
}

func (aC *ageCli) CommandConfigure(c *cli.Command) {
	aC.now = time.Now()

	c.SetName("date:dob").
		SetDescription("Calculate Date of Birth").
		AddArgument("name", "Your Name", true, arguments.String{
			Ptr:     &aC.name,
			MinRune: 5,
			Pattern: regexp.MustCompile(`^[A-Za-z]*$`),
		}).
		AddArgument("dob", "Your date of birth", true, arguments.Time{
			Ptr:    &aC.dob,
			Format: "02-01-2006",
			Max:    aC.now,
		})
}

func (aC *ageCli) CommandExecute() {
	duration := aC.now.Sub(aC.dob)
	fmt.Printf("Hello, my name is '%s' and I'm '%d' years old", aC.name,
		time.Date(0, 0, 0, 0, 0, int(duration.Seconds()), 0, aC.now.Location()).Year())
	fmt.Println()
}

func init() {
	cli.RegisterCommand(&ageCli{})
}