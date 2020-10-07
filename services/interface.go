package services

import (
	"github.com/labstack/echo/v4"
)

// Service aims to help autodiscover our avialable services
// TODO Finish service interface and implement
type Service interface {
	// GetHandlers returns a map of GET endpoint strings followed by the representative handler function via the echo HandlerFunc interface
	GetHandlers() map[string]echo.HandlerFunc
}
