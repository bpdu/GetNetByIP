package main

import (
	"fmt"
	"os"
)

const HELP_MESSAGE = `
Usage: go run GetNetByIP [OPTION]... IP_ADDRESS

Get network ID and some metadata by requested IP address.

Mandatory arguments to long options are mandatory for short options too.

Options:

  -i, --info                  Returns a network ID.

`

func main() {

	if len(os.Args) == 1 {
		fmt.Print(HELP_MESSAGE)
	}
}
