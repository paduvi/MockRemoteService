package controllers

import (
	"github.com/kataras/iris"
	IndexController "github.com/paduvi/MockRemoteService/controllers/index"
	MessageController "github.com/paduvi/MockRemoteService/controllers/message"
)

func WithRouter(app *iris.Application) {
	mainRouter := app.Party("/")

	IndexController.EquipRouter(mainRouter)
	MessageController.EquipRouter(mainRouter)
}
