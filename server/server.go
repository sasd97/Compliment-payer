package server

import (
	"github.com/go-martini/martini"
	"github.com/sasd97/flatterer/constants"
	"github.com/sasd97/flatterer/controllers"
)

func Make() *martini.ClassicMartini {
	server := martini.Classic()
	server.Get(constants.HOME, controllers.HandleHome)
	server.Get(constants.START, controllers.HandleStart)
	server.Get(constants.AUTHORIZE, controllers.HandleAuthorize)
	server.Get(constants.HANDLE_TOKEN, controllers.HandleCode)
	return server
}
