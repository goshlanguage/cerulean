package cerulean

import (
	"fmt"

	"github.com/goshlanguage/cerulean/services/subscriptions"
	"github.com/labstack/echo"
)

// Cerulean holds the handlers and port to instantiate the mock server
type Cerulean struct {
	// Addr is the address that Cerulean should listen at, eg: 127.0.0.1:51234
	Addr string
	// Handlers      map[string]http.Handler
	// Mux           *http.ServeMux
	Router    *echo.Echo
	Inventory Inventory
}

type Inventory struct {
	Subscriptions *[]subscriptions.Subscription
}

// New takes in a stringified address, eg: "127.0.0.1:8080" or ":8080",
//   as well as a mock subscriptionID to instiate your Cerulean instance with
//   and returns a the mock server
//
// New generates a local address to be passed in when initializing a `BaseClient`
//   in order to point it at the mock server.
func New(subscriptionID string) Cerulean {
	addr := ":0"
	// initSub is our initial SubscriptionID. This is important because there isn't an API route to create a SubscriptionID
	// (or if there is please open an issue and let us know!)
	// initSub := subscriptions.NewSubscription(subscriptionID)
	// subs := &[]subscriptions.Subscription{initSub}

	// TODO: Automatic iteration over handlers
	// handlers := make(map[string]http.Handler)
	// handlers["/subscriptions"] = subscriptions.GetSubscriptionsHandler(subs)
	// handlers["/subscriptions/{subscriptionID}/resourceGroups/{resourceGroupName}"] = resourceGroups.PutResourceGroupsHandler(subs)

	// mux := http.NewServeMux()
	// for route, handler := range handlers {
	// mux.Handle(route, handler)
	// }

	server := Cerulean{
		Addr: addr,
		// Handlers:      handlers,
		// Mux:           mux,
		Router: echo.New(),
		Inventory: Inventory{
			Subscriptions: &[]subscriptions.Subscription{
				subscriptions.NewSubscription(subscriptionID),
			},
		},
	}

	// server.ListenAndServe()
	// return server

	// Routes
	// TODO: Pass subscription inventory to the subscriptions handler
	server.Router.GET("/subscriptions", subscriptions.GetSubscriptionsHandler(server.Inventory.Subscriptions))
	// server.Router.PUT("/subscriptions/:subscriptionID/resourceGroups/:resourceGroupName", resourceGroups.PutResourceGroupsHandler)

	// Start server
	// listener, err := net.Listen("tcp", ":0")
	// if err != nil {
	// return err
	// }

	// server.Addr = fmt.Sprintf(":%v", listener.Addr().(*net.TCPAddr).Port)

	// go http.Serve(listener, server.Mux)
	go server.Router.Logger.Fatal(server.Router.Start(addr))

	server.Addr = server.Router.TLSServer.Addr

	return server
}

// GetBaseClientURI returns the address string in the form consumable by say an azure-sdk-for-go BaseClient
func (server *Cerulean) GetBaseClientURI() string {
	return fmt.Sprintf("http://127.0.0.1%s", server.Addr)
}
