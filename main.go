package main

import (
	"fmt"
	_ "github.com/CJ-Jackson/example-cli/src"
	"github.com/cjtoolkit/cli"
)

func main() {
	fmt.Println("CLI Example:")
	fmt.Println()

	cli.Run()
}
