package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

//go:embed index.gohtml
var echoTemplateStr string
var echoTemplate = template.Must(template.New("").Parse(echoTemplateStr))

//go:embed bootstrap.min.css
var bootstrap string

func main() {
	port := parsePort()

	log.Printf("Listening on http://localhost:%d", port)

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.Printf("%s %s", req.Method, req.URL.Path)

		env := map[string]string{}
		for _, value := range os.Environ() {
			pair := strings.Split(value, "=")
			if len(pair) != 2 {
				panic("no!")
			}
			env[pair[0]] = pair[1]
		}

		data := struct {
			Style  template.CSS
			Method string
			URL    *url.URL
			Header http.Header
			Env    map[string]string
		}{
			Style:  template.CSS(bootstrap),
			Method: req.Method,
			URL:    req.URL,
			Header: req.Header,
			Env:    env,
		}
		echoTemplate.Execute(res, data)
	})

	log.Panic(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func parsePort() int {
	if len(os.Args) != 2 {
		fmt.Println("Usage: echo <port>")
		os.Exit(1)
	}

	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf("Error parsing %q as port", os.Args[1])
	}

	return port
}
