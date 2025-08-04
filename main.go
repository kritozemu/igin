package main

import (
	"imitate_gin/igin"
	"net/http"
)

func main() {
	server := igin.NewEngine()

	server.GET("/", func(ctx *igin.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello World</h1>")
	})

	server.GET("/hello", func(ctx *igin.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})

	server.Run(":8080")
}
