package resourcegroups

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// PutHandler is the PUT method handler for /subscriptions/{subscription-id}/resourceGroups
func (svc *Service) PutHandler() echo.HandlerFunc {
	// subscriptionID string, resourceGroupName string) http.Handler {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, NewResourceGroupsResponse("subscriptionID", "resourceGroupName"))
	}
}