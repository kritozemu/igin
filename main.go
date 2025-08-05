package main

import (
	"imitate_gin/igin"
	"log"
	"net/http"
	"time"
)

func onlyForV2() igin.HandlerFunc {
	return func(ctx *igin.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		log.Printf("[%d] %s in %v for group v2", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
		ctx.Fail(500, "Internal Server Error")
		// Calculate resolution time

	}
}

func main() {
	r := igin.NewEngine()
	r.Use(igin.Logger()) // global midlleware
	r.GET("/", func(c *igin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	//v2.Use(onlyForV2()) // v2 group middleware

	v2.GET("/hello/:name", func(c *igin.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.Run(":8080")
}
