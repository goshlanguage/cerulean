package resourcegroups

import (
	"net/http"

	"github.com/goshlanguage/cerulean/internal/models"
	"github.com/labstack/echo/v4"
)

// GetHandler is the GET method handler for /subscriptions/{subscription-id}/resourceGroup/{resource-group-name}
func (svc *Service) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		return svc.getResourceGroup(c)
	}
}

func (svc *Service) getResourceGroup(c echo.Context) error {
	resourceGroup, err := svc.GetResourceGroup(c.Param("subscriptionID"), c.Param("resourceGroupName"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.NewCloudError(http.StatusBadRequest, err))
	}

	return c.JSON(http.StatusOK, resourceGroup)
}

// PutHandler is the PUT method handler for /subscriptions/{subscription-id}/resourceGroup/{resource-group-name}
func (svc *Service) PutHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		// TODO: Check if the resource group already exists in the store; return 201 if created and 200 if exists
		err := svc.AddResourceGroup(c.Param("subscriptionID"), c.Param("resourceGroupName"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, models.NewCloudError(http.StatusInternalServerError, err))
		}

		return svc.getResourceGroup(c)
	}
}
