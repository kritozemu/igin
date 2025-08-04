package main

import (
	"fmt"
	"imitate_gin/igin"
	"net/http"
)

func main() {
	server := igin.NewEngine()

	server.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	})

	server.GET("/hello", func(writer http.ResponseWriter, request *http.Request) {
		for k, v := range request.Header {
			fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
		}
	})

	server.Run(":8080")
}
