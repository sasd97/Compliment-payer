package services

import (
	"github.com/sasd97/flatterer/config"
	"github.com/yanple/vk_api"
)

var api vk_api.Api

func IsAuthorized() bool {
	return api.AccessToken != ""
}

func AccessToken() string {
	return api.AccessToken
}

func Authorize() (string, error) {
	return api.GetAuthUrl(config.REDIRECT_LINK, config.CLIENT_ID, config.SCOPE)
}

func OAuth(code string) error {
	return api.OAuth(config.REDIRECT_LINK, config.API_KEY, config.CLIENT_ID, code)
}
