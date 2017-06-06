package views

import (
	"net/http"

	"github.com/k8guard/k8guard-report/templates"
)

var (
	Version string
	Build   string
)

type context struct {
	Name    string
	Version string
	Build   string
}

func Index(w http.ResponseWriter, r *http.Request) {

	c := context{Name: "K8Guard", Version: Version, Build: Build}
	err := templates.IndexTemplate.Execute(w, c)
	if err != nil {
		panic(err)
	}

}
