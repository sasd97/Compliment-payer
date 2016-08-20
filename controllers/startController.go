package controllers

import (
	"fmt"
	"net/http"

	"github.com/sasd97/flatterer-vk/constants"
	"github.com/sasd97/flatterer-vk/services"
)

func HandleStart(w http.ResponseWriter, r *http.Request) {
	if services.IsAuthorized() {
		fmt.Fprint(w, services.AccessToken())
		services.Log()
		return
	}

	http.Redirect(w, r, constants.AUTHORIZE, http.StatusFound)
}
