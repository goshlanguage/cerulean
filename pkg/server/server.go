package server

import (
	"io"
	"log"
	"net/http"
)

// Server holds the handlers and port to instantiate the server
type Server struct {
	Handlers map[string]http.Handler
	Addr     string
}

// GetServer takes in a stringified address, eg: "127.0.0.1:8080" or ":8080", and returns a composed server
func GetServer(addr string) Server {
	handlers := make(map[string]http.Handler)
	handlers["/hello"] = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	})

	server := Server{
		Handlers: handlers,
		Addr:     addr,
	}

	return server
}

// ListenAndServe starts our server
func (server Server) ListenAndServe() {
	for route, handler := range server.Handlers {
		http.Handle(route, handler)
	}

	log.Fatal(http.ListenAndServe(server.Addr, nil))
}
