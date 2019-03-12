/*
Package routes is a small, no-frills routing library for fasthttp. It's designed to sit within a fasthttp
aware service to determine which route, from a map, to direct a ctx at.

It is designed:

1. To contain no third party module (beyond fasthttp)

2. To be as unobtrusive as possible

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


The above will return `Hello james` for GET /hello/james, and a 404 for anything else.
*/
package routes
