package views

import (
	"net/http"

	"github.com/k8guard/k8guard-report/db"
)

func Ready(w http.ResponseWriter, r *http.Request) {
	v := db.VActionResponseModel{}
	if err := v.Ping(); err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	w.WriteHeader(http.StatusOK)
}
