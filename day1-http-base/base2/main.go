package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engine struct{}

//只要实现了这个Serv
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q \n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT found: %s\n", req.URL)
	}
}
func main() {
	engine := new(Engine)
	log.Fatal(http.ListenAndServe(":9999", engine))
}
