package igin

import (
	"net/http"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
	router *router
}

func NewEngine() *Engine {
	return &Engine{router: newRouter()}
}

func (e *Engine) addRoute(method string, path string, hf HandlerFunc) {
	e.router.addRoute(method, path, hf)
}

func (e *Engine) GET(path string, hf HandlerFunc) {
	e.addRoute(http.MethodGet, path, hf)
}

func (e *Engine) POST(path string, hf HandlerFunc) {
	e.addRoute(http.MethodPost, path, hf)
}

func (e *Engine) DELETE(path string, hf HandlerFunc) {
	e.addRoute(http.MethodDelete, path, hf)
}

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := NewContext(w, req)
	e.router.handle(c)
}
