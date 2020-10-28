package resourcegroups

import (
	"net/http"

	"github.com/goshlanguage/cerulean/internal/services"
	"github.com/goshlanguage/cerulean/pkg/lightdb"
)

const serviceKey = "resourcegroups"

// Service satisfies the Service interface, and is used to start and maintain the ResourceGroup Service
type Service struct {
	Store *lightdb.Store
}

// NewService is a factory for Service, which satisfies the services.Service interface
func NewService(s *lightdb.Store) services.Service {
	service := &Service{
		Store: s,
	}

	return service
}

// GetHandlers returns the HTTP Echo handlers that the service needs in order to operate
func (svc *Service) GetHandlers() map[string]services.Handler {
	svcMap := make(map[string]services.Handler)
	svcMap["/subscriptions/:subscriptionId/resourcegroups/:resourceGroupName"] = services.Handler{
		Verb: http.MethodGet,
		Func: svc.GetHandler(),
	}
	svcMap["/subscriptions/:subscriptionId/resourcegroups/:resourceGroupName"] = services.Handler{
		Verb: http.MethodPut,
		Func: svc.PutHandler(),
	}
	return svcMap
}
