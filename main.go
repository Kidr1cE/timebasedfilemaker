package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
	"time"
)

const defaultTemplate = "{{.name}}-{{.now}}.txt"

func main() {
	conf, err := os.ReadFile("/usr/local/time-based-filename.tmpl")
	if err != nil {
		fmt.Println("Using default template")
		conf = []byte(defaultTemplate)
	}
	tmpl := template.Must(template.New("fileNameTemplate").Parse(string(conf)))
	name := os.Args[1]
	now := time.Now().Format("20060102")

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, map[string]string{"name": name, "now": now})

	_, err = os.Create(buf.String())
	if err != nil {
		fmt.Println(err)
		return
	}
}
