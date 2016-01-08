package main

import (
	"fmt"
	"net/http"
)

type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHello(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Kitty!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9000", mux)
}
