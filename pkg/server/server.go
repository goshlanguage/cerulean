package server

import (
	"log"
	"net/http"

	"github.com/goshlanguage/cerulean/pkg/services/subscriptions"
)

// Server holds the handlers and port to instantiate the server
type Server struct {
	Addr          string
	Handlers      map[string]http.Handler
	Subscriptions []*subscriptions.Subscription
}

// GetServer takes in a stringified address, eg: "127.0.0.1:8080" or ":8080", and returns a composed server
func GetServer(addr string) Server {
	handlers := make(map[string]http.Handler)
	// TODO: Automatic iteration over handlers
	subs := []*subscriptions.Subscription{}
	handlers["/subscriptions"] = http.HandlerFunc(subscriptions.GetSubscriptionsHandler(subs))

	server := Server{
		Addr:          addr,
		Handlers:      handlers,
		Subscriptions: subs,
	}

	return server
}

// ListenAndServe starts our server. We pass in our inventory to handlers to setup our routes
func (server Server) ListenAndServe() {
	for route, handler := range server.Handlers {
		http.Handle(route, handler)
	}

	log.Fatal(http.ListenAndServe(server.Addr, nil))
}
