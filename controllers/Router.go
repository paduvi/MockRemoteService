package controllers

import (
	"github.com/kataras/iris"
	. "github.com/paduvi/MockRemoteService/controllers/index"
	. "github.com/paduvi/MockRemoteService/controllers/message"
)

func WithRouter(app *iris.Application) {
	mainRouter := app.Party("/")

	EquipIndexRouter(mainRouter)
	EquipMessageRouter(mainRouter)
}
