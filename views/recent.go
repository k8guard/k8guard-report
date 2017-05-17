package views

import (
	"net/http"

	"github.com/k8guard/k8guard-report/db"
	"github.com/k8guard/k8guard-report/templates"

	libs "github.com/k8guard/k8guardlibs"
)

func Recent(w http.ResponseWriter, r *http.Request) {

	v := db.VActionResponseModel{}

	err := r.ParseForm()
	if err != nil {
		libs.Log.Debug("Parsing Form Error", err)

	}
	namespace := r.Form.Get("namespace")

	c := v.GetAllByNameSpace("default")
	if namespace != "" {
		// Getting all
		// c := v.GetAll()
		libs.Log.Debug("Parsed namspace from form is ", namespace)
		c = v.GetAllByNameSpace(namespace)
	}
	err = templates.RecentTemplate.Execute(w, c)
	if err != nil {
		panic(err)
	}

}
