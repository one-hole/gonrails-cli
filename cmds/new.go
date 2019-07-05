package cmds

import (
	"fmt"
	"log"
	"os"

	"github.com/one-hole/gonrails-cli/helper"
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
	touchControllers(projectName)
	touchSerializers(projectName)
}

type ventory struct {
	ModuleName string
}

func touchMain(moduleName string) {

	createFile(
		fmt.Sprintf("%s/%s/main.go", pwd, moduleName),
		fmt.Sprintf("%s/templates/main.go.template", helper.ProjectPath),
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
		fmt.Sprintf("%s/templates/config/config.go.template", helper.ProjectPath),
		nil,
	)

	createFile(
		fmt.Sprintf("%s/%s/config/app.yml", pwd, moduleName),
		fmt.Sprintf("%s/templates/config/app.yml", helper.ProjectPath),
		nil,
	)

	createFile(
		fmt.Sprintf("%s/%s/config/config.yml", pwd, moduleName),
		fmt.Sprintf("%s/templates/config/config.yml", helper.ProjectPath),
		nil,
	)
}

/*
	1. routes/base.go
	2. routes/admin/base.go
*/

func touchRouter(moduleName string) {
	createFile(
		fmt.Sprintf("%s/%s/routes/base.go", pwd, moduleName),
		fmt.Sprintf("%s/templates/routes/base.go.template", helper.ProjectPath),
		ventory{
			ModuleName: moduleName,
		},
	)

	_ = os.Mkdir(fmt.Sprintf("%s/routes/%s", moduleName, "admin"), os.ModePerm)

	createFile(
		fmt.Sprintf("%s/%s/routes/admin/base.go", pwd, moduleName),
		fmt.Sprintf("%s/templates/routes/admin/base.go.template", helper.ProjectPath),
		ventory{
			ModuleName: moduleName,
		},
	)
}

/*
	1. controllers/base.go
	2. controllers/admin/home/index.go
*/

func touchControllers(moduleName string) {
	createFile(
		fmt.Sprintf("%s/%s/controllers/base.go", pwd, moduleName),
		fmt.Sprintf("%s/templates/controllers/base.go.template", helper.ProjectPath),
		nil,
	)

	_ = os.Mkdir(fmt.Sprintf("%s/controllers/%s", moduleName, "home"), os.ModePerm)
	_ = os.Mkdir(fmt.Sprintf("%s/controllers/%s", moduleName, "books"), os.ModePerm)

	createFile(
		fmt.Sprintf("%s/%s/controllers/home/index.go", pwd, moduleName),
		fmt.Sprintf("%s/templates/controllers/home/index.go.template", helper.ProjectPath),
		nil,
	)

	createFile(
		fmt.Sprintf("%s/%s/controllers/books/index.go", pwd, moduleName),
		fmt.Sprintf("%s/templates/controllers/books/index.go", helper.ProjectPath),
		nil,
	)

	createFile(
		fmt.Sprintf("%s/%s/controllers/books/show.go", pwd, moduleName),
		fmt.Sprintf("%s/templates/controllers/books/show.go.template", helper.ProjectPath),
		ventory{
			ModuleName: moduleName,
		},
	)
}

/*
	Serializers
*/

func touchSerializers(moduleName string) {
	createFile(
		fmt.Sprintf("%s/%s/serializers/book_serializer.go", pwd, moduleName),
		fmt.Sprintf("%s/templates/serializers/book_serializer.go", helper.ProjectPath),
		nil,
	)
}
