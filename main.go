package main

import (
	"fmt"
	"net/http"

	"github.com/gocraft/web"
	"github.com/unrolled/render"
)

// Context for holding data between the request and the template
type Context struct {
	RawStruct    string
	ConvertedPHP string
}

var rend = render.New(render.Options{
	Extensions: []string{".html"},
	Directory:  "templates",
	Layout:     "index",
})

func main() {
	// Setup router
	rootRouter := web.New(Context{})

	rootRouter.Middleware(web.LoggerMiddleware)
	rootRouter.Middleware(web.StaticMiddleware("assets"))

	// Routes
	rootRouter.Get("/", (*Context).convertStruct)
	rootRouter.Post("/", (*Context).convertStruct)

	// Serve
	port := "3334"
	fmt.Printf("\x1b[32;1m --------- ConvertStructs [listening on port %s]\x1b[0m", port)
	http.ListenAndServe("localhost:"+port, rootRouter)
}
