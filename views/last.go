package views

import (
	"net/http"

	"strconv"

	"github.com/k8guard/k8guard-report/db"
	"github.com/k8guard/k8guard-report/templates"
	libs "github.com/k8guard/k8guardlibs"
)

func Last(w http.ResponseWriter, r *http.Request) {

	v := db.VActionResponseModel{}

	err := r.ParseForm()
	if err != nil {
		libs.Log.Debug("Parsing Form Error", err)

	}

	numberOfLast := r.Form.Get("numberOfRecentToQuery")
	i, err := strconv.Atoi(numberOfLast)
	libs.Log.Debug("the parsed number is ", i)
	if i < 1 {
		i = 10
	}
	c := v.GetLastActions(i)
	err = templates.AllTemplate.Execute(w, c)
	if err != nil {
		panic(err)
	}

}
