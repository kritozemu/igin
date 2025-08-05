package igin

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePath(t *testing.T) {
	ok := reflect.DeepEqual(parsePath("/hello/:name"), []string{"hello", ":name"})
	ok = ok && reflect.DeepEqual(parsePath("/assets/*"), []string{"assets", "*"})
	ok = ok && reflect.DeepEqual(parsePath("/assets/*filepath/*"), []string{"assets", "*filepath"})
	if !ok {
		t.Fatal("test parsePath fail")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/dora")
	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.path != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "dora" {
		t.Fatal("name should be dora")
	}

	fmt.Printf("matched path: %s\n,  params['name']: %s\n", n.path, ps["name"])
}
