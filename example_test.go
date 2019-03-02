package routes_test

import (
	"fmt"

	"github.com/jspc/routes"
	"github.com/valyala/fasthttp"
)

func ExampleRoutes_Add_simple() {
	r := routes.New()
	r.Add("/", func(ctx *fasthttp.RequestCtx) {
		fmt.Println("root")
	})

	req := fasthttp.AcquireRequest()
	req.SetRequestURI("/")
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()

	c := &fasthttp.RequestCtx{
		Request:  *req,
		Response: *resp,
	}

	r.Route(c)
	// Output: root
}

func ExampleRoutes_Add_missing() {
	r := routes.New()
	r.Add("/", func(ctx *fasthttp.RequestCtx) {
		fmt.Println("root")
	})

	r.Catcher = func(_ *fasthttp.RequestCtx) {
		fmt.Println("404")
	}

	req := fasthttp.AcquireRequest()
	req.SetRequestURI("/this/url/doesnt/exist")
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()

	c := &fasthttp.RequestCtx{
		Request:  *req,
		Response: *resp,
	}

	r.Route(c)
	// Output: 404
}

func ExampleRoutes_Add_params() {
	r := routes.New()
	r.Add("/hello/:name", func(ctx *fasthttp.RequestCtx) {
		name := ctx.UserValue("name")
		fmt.Printf("Pleased to meet you %v\n", name)
	})

	req := fasthttp.AcquireRequest()
	req.SetRequestURI("/hello/jspc")
	req.Header.SetMethod("GET")

	resp := fasthttp.AcquireResponse()

	c := &fasthttp.RequestCtx{
		Request:  *req,
		Response: *resp,
	}

	r.Route(c)
	// Output: Pleased to meet you jspc
}
