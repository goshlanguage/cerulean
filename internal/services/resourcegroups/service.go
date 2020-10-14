package resourcegroups

import (
	"net/http"

	"github.com/goshlanguage/cerulean/internal/services"
)

// Service satisfies the Service interface, and is used to start and maintain the ResourceGroup Service
type Service struct {
	// ResourceGroups is our statebag
	ResourceGroups []ResourceGroup
}

// NewService is a factory for Service, which satisfies the services.Service interface
func NewService() services.Service {
	return &Service{}
}

// GetHandlers returns the HTTP GET Echo handlers that the service needs in order to operate
func (svc *Service) GetHandlers() map[string]services.Handler {
	svcMap := make(map[string]services.Handler)
	svcMap["/subscriptions/:subscriptionId/resourcegroups/:resourceGroupName"] = services.Handler{http.MethodPut, svc.PutHandler()}
	return svcMap
}
