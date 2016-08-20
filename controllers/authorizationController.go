package controllers

import (
	"net/http"

	"github.com/sasd97/flatterer/constants"
	"github.com/sasd97/flatterer/services"
)

func HandleAuthorize(w http.ResponseWriter, r *http.Request) {
	authUrl, err := services.Authorize()

	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, authUrl, http.StatusFound)
}

func HandleCode(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	err := services.OAuth(code)

	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, constants.START, http.StatusFound)
}
