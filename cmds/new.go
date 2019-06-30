package cmds

import (
	"fmt"
	"log"
	"os"
)

// New - Creates a new gonrails project
func New(projectName string) {

	log.Printf("Creating gonrails project named %s ...", projectName)
	_ = os.Mkdir(projectName, os.ModePerm)

	for _, name := range dirs {
		log.Println(fmt.Sprintf("Making Dir: %s/%s", projectName, name))
		_ = os.Mkdir(fmt.Sprintf("%s/%s", projectName, name), os.ModePerm)
	}

	touchMain(projectName)
	touchConfig(projectName)
	touchRouter(projectName)

}

type ventory struct {
	ModuleName string
}

func touchMain(moduleName string) {

	createFile(
		fmt.Sprintf("%s/%s/main.go", pwd, moduleName),
		fmt.Sprintf("%s/templates/main.go.template", pwd),
		ventory{
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
	createFile(
		fmt.Sprintf("%s/%s/config/config.go", pwd, moduleName),
		fmt.Sprintf("%s/templates/config/config.go.template", pwd),
		nil,
	)

	createFile(
		fmt.Sprintf("%s/%s/config/app.yml", pwd, moduleName),
		fmt.Sprintf("%s/templates/config/app.yml", pwd),
		nil,
	)

	createFile(
		fmt.Sprintf("%s/%s/config/config.yml", pwd, moduleName),
		fmt.Sprintf("%s/templates/config/config.yml", pwd),
		nil,
	)
}

func touchRouter(moduleName string) {
	createFile(
		fmt.Sprintf("%s/%s/routes/base.go", pwd, moduleName),
		fmt.Sprintf("%s/templates/routes/base.go.template", pwd),
		ventory{
			ModuleName: moduleName,
		},
	)

	_ = os.Mkdir(fmt.Sprintf("%s/routes/%s", moduleName, "admin"), os.ModePerm)
}
