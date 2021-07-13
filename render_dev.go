//+build dev

package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
)

// In dev we hot-reload our template
func render(rw http.ResponseWriter, data interface{}) error {
	tmplBytes, err := os.ReadFile("index.gohtml")
	if err != nil {
		renderError(rw, err)
		return err
	}

	tmpl, err := template.New("").Parse(string(tmplBytes))
	if err != nil {
		renderError(rw, err)
		return err
	}

	buf := strings.Builder{}

	if err := tmpl.Execute(&buf, data); err != nil {
		renderError(rw, err)
		return err
	}

	if _, err := io.Copy(rw, strings.NewReader(buf.String())); err != nil {
		panic(err)
	}

	return nil
}

func renderError(rw http.ResponseWriter, err error) {
	rw.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(rw, "Error: %s", err.Error())
}
