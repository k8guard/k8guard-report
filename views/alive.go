package views

import (
	"net/http"
)

func Alive(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
