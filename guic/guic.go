package guic

import (
	"net/http"

	"github.com/deweppro/go-static"
)

//go:generate static ./../dist/solocmsadminui ui

var ui static.Reader

func ResponseWrite(w http.ResponseWriter, r *http.Request) error {
	filename := r.RequestURI
	switch filename {
	case "", "/":
		filename = "/index.html"
	default:
	}

	return ui.ResponseWrite(w, filename)
}
