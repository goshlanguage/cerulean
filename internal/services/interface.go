package services

import (
	"github.com/labstack/echo/v4"
)

// Handler is a helper struct that allows us to iterate over handlers to pass to Echo
type Handler struct {
	// Verb is a string containing an HTTP verb value (https://golang.org/src/net/http/method.go)
	Verb string
	Func echo.HandlerFunc
}

// Service aims to help autodiscover our available services
type Service interface {
	// GetHandlers returns a map of key/value pairs where the key is the HTTP route and the value is a Handler struct (an Echo handler function along with its HTTP verb)
	GetHandlers() map[string]Handler
}
