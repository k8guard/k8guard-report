package templates

import (
	"html/template"
	"io/ioutil"
	"strings"
)

var templates *template.Template
var IndexTemplate *template.Template
var LastTemplate *template.Template
var RecentTemplate *template.Template

func PopulateTemplates() {
	allFiles := []string{}
	templatesDir := "./templates/"
	files, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename, ".html") {
			allFiles = append(allFiles, templatesDir+filename)
		}
	}

	templates, err = template.ParseFiles(allFiles...)
	if err != nil {
		panic(err)
	}
	IndexTemplate = templates.Lookup("index.html")
	LastTemplate = templates.Lookup("last.html")
	RecentTemplate = templates.Lookup("recent.html")
}
