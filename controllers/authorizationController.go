package controllers

import (
	"fmt"
	"net/http"

	"github.com/sasd97/flatterer-vk/constants"
	"github.com/sasd97/flatterer-vk/services"
)

func HandleAuthorize(w http.ResponseWriter, r *http.Request) {
	authUrl := services.Authorize()

	fmt.Println(authUrl)
	http.Redirect(w, r, authUrl, http.StatusFound)
}

func HandleSetup(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("access_token")
	expires := r.URL.Query().Get("expires_in")
	user := r.URL.Query().Get("user_id")

	services.Setup(token, user, expires)
	http.Redirect(w, r, constants.START, http.StatusFound)
}
