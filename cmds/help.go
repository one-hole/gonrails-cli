package cmds

import (
	"github.com/one-hole/gonrails-cli/helper"
)

// Help - gonrails-cli help

var useageTemplate = `{{ "gonrails-cli" | bold }} is the command line tools for your Gonrails Web Application.

{{"USEAGE" | bold | red}}

	gonrails-cli command [arguments]

{{"COMMANDS" | bold | red}}
	
	new       Create a new project
	help      Show the usage of commands
	generate  Generate code by commands

{{"EXAMPLES" | bold | red}}

	new

		'gonrails-cli new kalista'

	generate

		'gonrails-cli generate model user'
		'gonrails-cli generate controller users actions:['index', 'create', 'delete', 'show', 'update']'
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
