package cmds

import (
	"fmt"
	"log"
	"os"
)

// New - Creates a new gonrails project
func New(projectName string) {

	log.Printf("creating gonrails project named %s ...", projectName)
	os.Mkdir(projectName, os.ModePerm)

	for _, name := range dirs {
		log.Println(fmt.Sprintf("making dir: %s/%s", projectName, name))
		os.Mkdir(fmt.Sprintf("%s/%s", projectName, name), os.ModePerm)
	}

	touchMain(projectName)
}

func touchMain(modulename string) {
	type inventory struct {
		Modulename string
	}

	createFile(
		"/Users/ctao/Mission/Golang/src/github.com/one-hole/gonrails-cli/watermelon/main.go",
		"/Users/ctao/Mission/Golang/src/github.com/one-hole/gonrails-cli/templates/main.go.template",
		inventory{
			Modulename: modulename,
		},
	)
}
