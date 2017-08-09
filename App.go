package main

import (
	"github.com/kataras/iris"
	"github.com/paduvi/MockRemoteService/middlewares"
	"github.com/paduvi/MockRemoteService/controllers"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

func init()  {
	_ = godotenv.Load("deploy.env")
}

func main() {
	app := iris.New()

	app.OnErrorCode(iris.StatusInternalServerError, middlewares.ErrorHandler)
	if os.Getenv("LogLevel") == "DEBUG" {
		app.Use(middlewares.Logger)
	}

	controllers.WithRouter(app)

	// Listen for incoming HTTP/1.x & HTTP/2 clients on localhost port 8080.
	app.Run(iris.Addr(os.Getenv("Address")), iris.WithCharset("UTF-8"))
}
