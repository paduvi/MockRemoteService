package todo

import (
	"encoding/json"
	"github.com/kataras/iris/context"
	. "github.com/paduvi/MockRemoteService/models"
	"github.com/paduvi/MockRemoteService/actions"
	"io/ioutil"
	"io"
	"github.com/kataras/iris"
)

func MessageIndex(ctx context.Context) {
	ctx.ContentType("application/json")
	ctx.StatusCode(iris.StatusOK)
	done := make(chan Messages)
	go actions.ListMessage(done)

	if _, err := ctx.JSON(<-done); err != nil {
		panic(err)
	}
}

func MessageShow(ctx context.Context) {
	messageId, err := ctx.Params().GetInt("messageId")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	done := make(chan Message)
	go actions.FindMessage(messageId, done)
	if _, err := ctx.JSON(<-done); err != nil {
		panic(err)
	}
}

func MessageCreate(ctx context.Context) {
	var todo Message
	body, err := ioutil.ReadAll(io.LimitReader(ctx.Request().Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := ctx.Request().Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		ctx.ContentType("application/json")
		ctx.StatusCode(iris.StatusUnprocessableEntity)
		if _, err := ctx.JSON(err); err != nil {
			panic(err)
		}
	}

	done := make(chan Message)
	go actions.CreateMessage(todo, done)
	ctx.ContentType("application/json")
	ctx.StatusCode(iris.StatusCreated)
	if _, err := ctx.JSON(<-done); err != nil {
		panic(err)
	}
}

func MessageDelete(ctx context.Context) {
	messageId, err := ctx.Params().GetInt("messageId")
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		return
	}
	done := make(chan error)
	go actions.DestroyMessage(messageId, done)
	if err := <-done; err != nil {
		panic(err)
	}
	ctx.Text("Destroy Message Successfully.")
}
