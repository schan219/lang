package main

import (
	"fmt"
	"lang/internals/cli"
)

func main() {
	fmt.Println("DEBUG: starting..")
	conf = cli.Start()
	// Pass in the configuration to the reader.


	fmt.Println("DEBUG: ending..")
}
