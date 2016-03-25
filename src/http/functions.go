package http

import (
	"fmt"
	"github.com/CJ-Jackson/example-cli/src/global"
	"net/http"
)

func SimpleHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello World</h1>")

	if global.GetGlobal().Prod {
		fmt.Fprintln(w, "<p>I am in production mode</p>")
	} else {
		fmt.Fprintln(w, "<p>I am in development mode</p>")
	}
}
