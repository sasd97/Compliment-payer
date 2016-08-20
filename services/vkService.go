package services

import (
	"fmt"

	"github.com/sasd97/flatterer-vk/config"
)

type API struct {
	AccessToken string
	UserId      string
	ExpiresIn   string
}

var api API

func IsAuthorized() bool {
	return api.AccessToken != ""
}

func AccessToken() string {
	return api.AccessToken
}

func Authorize() string {
	return fmt.Sprintf("https://oauth.vk.com/authorize?client_id=%s&redirect_uri=%s&response_type=token&scope=messages",
		config.CLIENT_ID,
		config.REDIRECT_SLINK)
}

func Setup(accessToken, userId, expiresIn string) {
	api.AccessToken = accessToken
	api.UserId = userId
	api.ExpiresIn = expiresIn
}

func Log() {
	fmt.Printf("Access token: %s\nUser id: %s\nExpires in: %s\n", api.AccessToken, api.UserId, api.ExpiresIn)
}
