package igin

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func NewEngine() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (e *Engine) addRoute(method string, path string, hf HandlerFunc) {
	key := method + "-" + path
	e.router[key] = hf
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
	key := req.Method + "-" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
