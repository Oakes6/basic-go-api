package main

import (
	"html"
	"io"
	"log"
	"net/http"
)

func mirror(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, html.EscapeString(r.URL.Path))
}

var mux map[string]func(w http.ResponseWriter, r *http.Request)

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: &myHandler{},
	}

	mux = make(map[string]func(w http.ResponseWriter, r *http.Request))
	mux["/"] = mirror

	log.Fatal(server.ListenAndServe())
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	io.WriteString(w, "My server: "+r.URL.String())
}
