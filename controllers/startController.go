package controllers

import (
	"fmt"
	"net/http"

	"github.com/sasd97/flatterer/constants"
	"github.com/sasd97/flatterer/services"
)

func HandleStart(w http.ResponseWriter, r *http.Request) {
	if services.IsAuthorized() {
		fmt.Fprint(w, services.AccessToken())
		return
	}

	http.Redirect(w, r, constants.AUTHORIZE, http.StatusFound)
}
