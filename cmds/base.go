package cmds

import (
	"io/ioutil"
	"os"
	"strings"
	"text/template"

	"github.com/one-hole/gonrails-cli/helper"
)

var (
	dirs = []string{"config", "routes", "models", "controllers", "serializers", "helper", "workers"}
)

// create file with
func createFile(filePath string, templatePath string, data interface{}) {
	file, err := os.Create(filePath)
	helper.PanicError(err)
	contents, _ := ioutil.ReadFile(templatePath)
	result := strings.Replace(string(contents), "\n", "", 1)
	var t = template.Must(template.New("temp").Parse(result))
	err = t.Execute(file, data)
	helper.PanicError(err)
}
