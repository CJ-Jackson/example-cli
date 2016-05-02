package http

import (
	"fmt"
	"github.com/cjtoolkit/cli"
	"github.com/cjtoolkit/cli/options"
	"net/http"
)

type httpCommand struct {
	address string
}

func (hC *httpCommand) CommandConfigure(c *cli.Command) {
	hC.address = ":8080"

	c.SetName("http:start:server").
		SetDescription("Run HTTP Server").
		AddOption("address", "Listening address", options.NewString(&hC.address))
}

func (hC *httpCommand) CommandExecute() {
	fmt.Printf("Running HTTP Server on '%s' (Ctrl + C to exit)...", hC.address)
	fmt.Println()

	http.ListenAndServe(hC.address, http.HandlerFunc(SimpleHelloWorld))
}

func init() {
	cli.RegisterCommand(&httpCommand{})
}
