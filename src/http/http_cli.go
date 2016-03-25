package http

import (
	"fmt"
	"github.com/cjtoolkit/cli"
	"github.com/cjtoolkit/cli/arguments"
	"net/http"
)

type httpCli struct {
	address string
}

func (hC *httpCli) CommandConfigure(c *cli.Command) {
	hC.address = ":8080"

	c.SetName("http:start:server").
		SetDescription("Run HTTP Server").
		AddArgument("address", "Listening address", false, arguments.String{Ptr: &hC.address})
}

func (hC *httpCli) CommandExecute() {
	fmt.Printf("Running HTTP Server on '%s' (Ctrl + C to exit)...", hC.address)
	fmt.Println()

	http.ListenAndServe(hC.address, http.HandlerFunc(SimpleHelloWorld))
}

func init() {
	cli.RegisterCommand(&httpCli{})
}