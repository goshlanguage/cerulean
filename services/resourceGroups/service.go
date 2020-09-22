package resourceGroups

import (
	"net/http"

	"github.com/goshlanguage/cerulean/services"
)

// ResourceGroupsService satisfies the Service interface, and is used to start and maintain the ResourceGroup Service
type ResourceGroupsService struct {
	// ResourceGroups is our statebag
	ResourceGroups []string
}

// NewResourceGroupsService is a factory for the ResourceGroupsService, which satisfies the services.Service interface
func NewResourceGroupsService() services.Service {
	return &ResourceGroupsService{}
}

// GetHandlers returns the HTTP GET Echo handlers that the service needs in order to operate
func (svc *ResourceGroupsService) GetHandlers() map[string]services.Handler {
	svcMap := make(map[string]services.Handler)
	svcMap["/providers/GroupsClient.CreateOrUpdate"] = services.Handler{http.MethodPut, svc.PutResourceGroupsHandler()}
	return svcMap
}
