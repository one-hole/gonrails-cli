package helper

import (
	"os"
	"text/template"
)

func Tmpl(text string, data interface{}) {
	output := os.Stdout
	t := template.New("HowTo")
	template.Must(t.Parse(text))
	err := t.Execute(output, data)
	PanicError(err)
}
