package resourceGroups

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TestPutResourceGroupsHandler sets up a server and tests the endpoint directly
func TestPutResourceGroupsHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/providers/GroupsClient.CreateOrUpdate", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	resourceGroupsService := ResourceGroupsService{}
	getHandler := resourceGroupsService.PutResourceGroupsHandler()

	// Assertions
	if assert.NoError(t, getHandler(ctx)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		// TODO: Maybe test that a UUID is returned
		assert.Contains(t, rec.Body.String(), "{\"subscriptionLink\":\"/subscriptions/")
	}
}
