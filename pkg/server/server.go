package server

import (
	"log"
	"net/http"

	"github.com/goshlanguage/cerulean/pkg/inventory"
	"github.com/goshlanguage/cerulean/pkg/services/subscriptions"
)

// Server holds the handlers and port to instantiate the server
type Server struct {
	Handlers  map[string]http.Handler
	Inventory *inventory.Inventory
	Addr      string
}

// GetServer takes in a stringified address, eg: "127.0.0.1:8080" or ":8080", and returns a composed server
func GetServer(addr string) Server {
	stateBag := &inventory.Inventory{}
	handlers := make(map[string]http.Handler)
	// TODO: Automatic iteration over handlers
	handlers["/subscriptions"] = subscriptions.GetSubscriptionsHandler("/subscriptions", stateBag)

	server := Server{
		Handlers:  handlers,
		Inventory: stateBag,
		Addr:      addr,
	}

	return server
}

// ListenAndServe starts our server. We pass in our inventory to handlers to setup our routes
func (server Server) ListenAndServe() {
	for route, handler := range server.Handlers {
		http.Handle(route, handler(server.Inventory))
	}

	log.Fatal(http.ListenAndServe(server.Addr, nil))
}
