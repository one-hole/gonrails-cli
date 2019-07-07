package cmds

import (
	"github.com/one-hole/gonrails-cli/helper"
)

// Help - gonrails-cli help

var useageTemplate = `gonrails-cli is the command line tools your Gonrails Web Application.

USEAGE

	gonrails-cli command [arguments]

COMMANDS
	
	new      -
	help     -
	generate -
`

func Help() {
}

func Useage() {
	helper.Tmpl(useageTemplate, nil)
}

/*
 1. gonrails-cli
 2. gonrails-cli help
 3. gonrails-cli help new
 4. gonrails-cli help generate
*/
