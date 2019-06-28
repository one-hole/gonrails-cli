package cmds

import (
	"fmt"
	"log"
	"os"
)

// New - Creates a new gonrails project
func New(projectName string) {

	log.Printf("creating gonrails project named %s ...", projectName)
	_ = os.Mkdir(projectName, os.ModePerm)

	for _, name := range dirs {
		log.Println(fmt.Sprintf("making dir: %s/%s", projectName, name))
		_ = os.Mkdir(fmt.Sprintf("%s/%s", projectName, name), os.ModePerm)
	}

	touchMain(projectName)
}

func touchMain(moduleName string) {

	log.Println("touching main.go ...")

	type inventory struct {
		ModuleName string
	}

	createFile(
		fmt.Sprintf("%s/%s/main.go", pwd, moduleName),
		fmt.Sprintf("%s/templates/main.go.template", pwd),
		inventory{
			ModuleName: moduleName,
		},
	)
}

/*
	1. config/config.go
	2. config/config.yml
	3. config/app.yml
	4. README.md
*/
func touchConfig(moduleName string) {

}

func touchRouter(moduleName string) {

}
