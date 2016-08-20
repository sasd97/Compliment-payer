package server

import (
	"github.com/go-martini/martini"
	"github.com/sasd97/flatterer-vk/constants"
	"github.com/sasd97/flatterer-vk/controllers"
)

func Make() *martini.ClassicMartini {
	server := martini.Classic()
	server.Get(constants.HOME, controllers.HandleHome)
	server.Get(constants.START, controllers.HandleStart)
	server.Get(constants.AUTHORIZE, controllers.HandleAuthorize)
	server.Get(constants.SETUP, controllers.HandleSetup)
	return server
}
