package igin

import (
	"log"
	"net/http"
)

type router struct {
	Handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		Handlers: make(map[string]HandlerFunc),
	}
}

func (r *router) addRoute(method string, path string, h HandlerFunc) {
	log.Printf("addRoute method:%s path:%s", method, path)
	key := method + "-" + path
	r.Handlers[key] = h
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.Handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 page not found:%s\n", c.Path)
	}
}
