package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var (
		flTemplate       = flag.String("template", "", "path to template")
		flJSON           = flag.String("json", "", "json input to use in template")
		flTemplateString = flag.Bool("string", false, "template is a string, not a file")
	)
	flag.Parse()

	var dict map[string]interface{}
	if err := json.Unmarshal([]byte(*flJSON), &dict); err != nil {
		log.Fatal(err)
	}

	var (
		err  error
		tmpl *template.Template
		name string
	)

	if *flTemplateString {
		name = "template"
		tmpl, err = template.New(name).Parse(*flTemplate)
	} else {
		name = filepath.Base(*flTemplate)
		tmpl, err = template.ParseFiles(*flTemplate)
	}
	if err != nil {
		log.Fatal(err)
	}

	if err := tmpl.ExecuteTemplate(os.Stdout, name, dict); err != nil {
		log.Fatal(err)
	}
}
