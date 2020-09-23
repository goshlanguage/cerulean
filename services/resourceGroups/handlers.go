package resourceGroups

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// PutResourceGroupsHandler is the PUT method handler for /subscriptions/{subscription-id}/resourceGroups
func (svc *ResourceGroupsService) PutResourceGroupsHandler() echo.HandlerFunc {
	// subscriptionID string, resourceGroupName string) http.Handler {
	return func(c echo.Context) error {
		b, err := json.Marshal(NewResourceGroupsResponse("subscriptionID", "resourceGroupName"))
		if err != nil {
			panic(err)
		}
		c.Response().Header().Set("Content-Type", "application/json")
		c.Response().Header().Set("Charset", "UTF-8")
		return c.JSON(http.StatusOK, b)
	}
}
