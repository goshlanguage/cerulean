package resourcegroups

import (
	"encoding/json"
	"errors"

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

// GetServiceHandlers returns the HTTP Echo handlers that the service needs in order to operate
func (svc *Service) GetServiceHandlers(e *echo.Echo) []*echo.Route {
	return []*echo.Route{
		e.GET("/subscriptions/:subscriptionID/resourcegroups/:resourceGroupName", svc.GetHandler()),
		e.PUT("/subscriptions/:subscriptionID/resourcegroups/:resourceGroupName", svc.PutHandler()),
	}
}

// GetResourceGroup returns a resource group for a specific subscription found in the Store
func (svc *Service) GetResourceGroup(subscriptionID string, resourceGroupName string) (ResourceGroup, error) {
	var resourceGroup ResourceGroup
	resourceGroupString := svc.Store.Get(serviceKey + "/" + subscriptionID + "/" + resourceGroupName)

	err := json.Unmarshal([]byte(resourceGroupString), &resourceGroup)
	if err != nil {
		return ResourceGroup{}, err
	}

	return resourceGroup, nil
}

// AddResourceGroup takes a resource group and adds it to the Store
func (svc *Service) AddResourceGroup(subscriptionID string, resourceGroupName string) error {
	if len(subscriptionID) == 0 {
		return errors.New("subscriptionID URI param cannot be empty")
	}

	if len(resourceGroupName) == 0 {
		return errors.New("resourceGroupName URI param cannot be empty")
	}

	resourceGroupsBytes, err := json.Marshal(NewResourceGroupsResponse(subscriptionID, resourceGroupName))
	if err != nil {
		return err
	}
	resourceGroupString := string(resourceGroupsBytes)

	svc.Store.Put(serviceKey+"/"+subscriptionID+"/"+resourceGroupName, resourceGroupString)

	return nil
}
