package cerulean

import (
	"fmt"
	"net"
	"net/http"

	"github.com/goshlanguage/cerulean/internal/services"
	"github.com/goshlanguage/cerulean/internal/services/subscriptions"
	"github.com/goshlanguage/cerulean/pkg/lightdb"
	"github.com/labstack/echo/v4"
)

// Cerulean holds the handlers and port to instantiate the mock server
type Cerulean struct {
	// Addr is the address that Cerulean should listen at, eg: 127.0.0.1:51234
	Addr string
	// BaseSubscriptionID is the base subscriptionID created and the default subscriptionID used for tests
	BaseSubscriptionID string
	Echo               *echo.Echo
	Services           []services.Service
	Store              *lightdb.Store
}

// New sets up an instance of our mock and returns it
//   and returns a the mock server
//
// New generates a local address to be passed in when initializing a `BaseClient`
//   in order to point it at the mock server.
func New() Cerulean {
	e := echo.New()
	s := lightdb.NewStore()
	e.HideBanner = true // Make log output less noisy by removing ASCII artwork

	subscriptionsSVC := subscriptions.NewService(s)
	baseSub := subscriptionsSVC.GetBaseSubscriptionID()

	svcs := []services.Service{
		subscriptionsSVC,
	}

	for _, service := range svcs {
		for endpoint, handlerStruct := range service.GetHandlers() {
			switch verb := handlerStruct.Verb; verb {
			case http.MethodGet:
				e.GET(endpoint, handlerStruct.Func)
			case http.MethodHead:
				e.HEAD(endpoint, handlerStruct.Func)
			case http.MethodPost:
				e.POST(endpoint, handlerStruct.Func)
			case http.MethodPut:
				e.PUT(endpoint, handlerStruct.Func)
			case http.MethodPatch:
				e.PATCH(endpoint, handlerStruct.Func)
			case http.MethodDelete:
				e.DELETE(endpoint, handlerStruct.Func)
			case http.MethodConnect:
				e.CONNECT(endpoint, handlerStruct.Func)
			case http.MethodOptions:
				e.OPTIONS(endpoint, handlerStruct.Func)
			case http.MethodTrace:
				e.TRACE(endpoint, handlerStruct.Func)
			}
		}
	}

	server := Cerulean{
		Addr:               ":0",
		BaseSubscriptionID: baseSub,
		Echo:               e,
		Services:           svcs,
		Store:              lightdb.NewStore(),
	}
	server.ListenAndServe()

	return server
}

// ListenAndServe starts our server.
func (server *Cerulean) ListenAndServe() error {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return err
	}
	server.Addr = fmt.Sprintf(":%v", listener.Addr().(*net.TCPAddr).Port)

	server.Echo.Listener = listener
	go server.Echo.Start(":0")

	return nil
}

// GetBaseClientURI returns the address string in the form consumable by say an azure-sdk-for-go BaseClient
func (server *Cerulean) GetBaseClientURI() string {
	return fmt.Sprintf("http://127.0.0.1%s", server.Addr)
}
