package main

import (
	"flag"
	"os"

	"github.com/one-hole/gonrails-cli/cmds"
	"github.com/one-hole/gonrails-cli/helper"
)

func main() {

	flag.Parse()
	args := flag.Args()

	if len(args) <= 1 {
		cmds.Useage()
		os.Exit(2)
		return
	}

	if "new" == args[0] {
		cmds.New(args[1])
	}

	if "generate" == args[0] {

	}

	helper.Tmpl("myabc", nil)
}
