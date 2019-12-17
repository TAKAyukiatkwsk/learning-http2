package main

import (
	"io"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	var srv http.Server
	srv.Addr = ":8080"

	http2.ConfigureServer(&srv, nil)
	http.HandleFunc("/", index)

	srv.ListenAndServeTLS("server.crt", "server.key")
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World! by HTTP/2")
}
