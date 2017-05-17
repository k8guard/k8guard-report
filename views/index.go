package views

import (
	"github.com/k8guard/k8guard-report/templates"
	"net/http"
	"time"
)

type context struct {
	Name string
	Time time.Time
}

func Index(w http.ResponseWriter, r *http.Request) {

	c := context{Name: "K8Guard", Time: time.Now()}
	err := templates.IndexTemplate.Execute(w, c)
	if err != nil {
		panic(err)
	}

}
