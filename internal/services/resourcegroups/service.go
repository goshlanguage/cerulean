package resourcegroups

import (
	"github.com/goshlanguage/cerulean/pkg/lightdb"
	"github.com/labstack/echo/v4"
)

const serviceKey = "resourcegroups"

// Service satisfies the Service interface, and is used to start and maintain the ResourceGroup Service
type Service struct {
	Store *lightdb.Store
}

// NewService is a factory for Service, which satisfies the services.Service interface
func NewService(s *lightdb.Store) *Service {
	service := &Service{
		Store: s,
	}

	return service
}

// GetAllHandlers returns the HTTP Echo handlers that the service needs in order to operate
func (svc *Service) GetAllHandlers(e *echo.Echo) []*echo.Route {
	return []*echo.Route{
		e.GET("/subscriptions/:subscriptionId/resourcegroups/:resourceGroupName", svc.GetHandler()),
		e.PUT("/subscriptions/:subscriptionId/resourcegroups/:resourceGroupName", svc.PutHandler()),
	}
}
