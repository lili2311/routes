package main

import (
	"fmt"

	"github.com/jspc/routes"
	"github.com/valyala/fasthttp"
)

type API struct {
	Message string
}

func (a API) Hello(ctx *fasthttp.RequestCtx) {
	name := ctx.UserValue("name").(string)

	fmt.Fprintf(ctx, "%s %s", a.Message, name)
}

func (a API) FourOhFour(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusNotFound)

	fmt.Fprintf(ctx, "%s not found", string(ctx.Path()))
}

func (a API) Handle(ctx *fasthttp.RequestCtx) {
	r := routes.New()
	r.Catcher = a.FourOhFour

	r.Add("/hello/:name", a.Hello)

	r.Route(ctx)
}

func main() {
	panic(fasthttp.ListenAndServe(":8080", API{"Hello"}.Handle))
}
