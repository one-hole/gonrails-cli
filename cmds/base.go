package cmds

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/one-hole/gonrails-cli/helper"
)

var (
	dirs   = []string{"config", "routes", "models", "controllers", "serializers", "helper", "workers"}
	pwd, _ = os.Getwd()
)

// create file with path and template
func createFile(filePath string, templatePath string, data interface{}) {

	log.Printf("Creating file : %s", filePath)

	file, err := os.Create(filePath)
	helper.PanicError(err)
	contents, _ := ioutil.ReadFile(templatePath)
	result := strings.Replace(string(contents), "\n", "", 1)
	var t = template.Must(template.New("main").Parse(result))
	err = t.Execute(file, data)
	helper.PanicError(err)
}

func mkdir() {

}
