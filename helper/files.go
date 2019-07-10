package helper

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
)

// ProjectPath for template files
var ProjectPath string

// Pwd for current dir
var Pwd, _ = os.Getwd()

func init() {
	if "true" == os.Getenv("DEV") {
		ProjectPath, _ = os.Getwd()
	} else {
		ProjectPath = fmt.Sprintf("%s/src/github.com/one-hole/gonrails-cli", os.Getenv("GOPATH"))
	}
}

// FileExists - file exists or not
func FileExists(path string) bool {
	_, err := os.Stat(path)
	if nil == err {
		return true
	}
	return false
}

// CreateFile -
func CreateFile(filePath string, templatePath string, data interface{}) {

	log.Printf("Creating file : %s", filePath)

	file, err := os.Create(filePath)
	PanicError(err)
	contents, _ := ioutil.ReadFile(templatePath)
	result := strings.Replace(string(contents), "\n", "", 1)
	var t = template.Must(template.New(filePath).Funcs(FuncMap()).Parse(result))
	err = t.Execute(file, data)
	PanicError(err)
}

// FindOrCreateDir - should be nested admin/user
func FindOrCreateDir(dirPath string) error {
	if FileExists(dirPath) {
		return nil
	}
	log.Printf("Making dir : %s\n", dirPath)
	err := os.MkdirAll(dirPath, os.ModePerm)
	return err
}
