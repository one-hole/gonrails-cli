package main

import (
	"flag"
	"os"

	"github.com/one-hole/gonrails-cli/cmds"
)

func main() {

	flag.Parse()
	args := flag.Args()

	if len(args) <= 1 {
		os.Exit(1)
		return
	}

	if "new" == args[0] {
		cmds.New(args[1])
	}

	if "generate" == args[0] {

	}
}
