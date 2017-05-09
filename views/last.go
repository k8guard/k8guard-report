package views

import (
	"k8guard-report/db"
	"k8guard-report/templates"
	"net/http"
)

func Last(w http.ResponseWriter, r *http.Request) {

	v := db.VActionResponseModel{}
	// Getting all
	c, err := v.GetLastAction()
	err = templates.LastTemplate.Execute(w, c)
	if err != nil {
		panic(err)
	}

}
