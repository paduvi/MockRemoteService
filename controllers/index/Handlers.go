package todo

import (
	"github.com/kataras/iris/context"
)

func Index(ctx context.Context) {
	ctx.Text("Ping!")
}

func SubIndex(ctx context.Context) {
	name := ctx.Params().Get("name")
	ctx.Text("Welcome, " + name + "!")
}
