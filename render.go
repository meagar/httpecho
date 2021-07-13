//+build release

package main

import (
	_ "embed"
	"html/template"
	"net/http"
)

//go:embed index.gohtml
var echoTemplateStr string
var echoTemplate = template.Must(template.New("").Parse(echoTemplateStr))

// In production we embed our template with the `embed` package
func render(rw http.ResponseWriter, data interface{}) error {
	return echoTemplate.Execute(rw, data)
}
