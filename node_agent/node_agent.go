package main

import (
	"os"

	cli "github.com/spiffe/node-agent/cli"
)

func main() {
	os.Exit(cli.Run(os.Args[1:]))
}