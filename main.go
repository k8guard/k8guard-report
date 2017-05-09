package main

import (
	"k8guard-report/db"
	"k8guard-report/templates"

	libs "github.com/k8guard/k8guardlibs"
)

func main() {
	templates.PopulateTemplates()
	err := db.Connect(libs.Cfg.CassandraHosts)
	if err != nil {
		panic(err.Error())
	}

	start_http_router()
}
