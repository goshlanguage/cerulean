package resourcegroups

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goshlanguage/cerulean/pkg/lightdb"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TestResourceGroupPutAndGet sets up a server and tests the endpoint directly
func TestResourceGroupPutAndGet(t *testing.T) {
	subscriptionID := "12345"
	resourceGroupName := "testResourceGroupName"

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/subscriptions/"+subscriptionID+"/resourcegroups/"+resourceGroupName, nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	s := lightdb.NewStore()
	resourceGroupsService := Service{
		Store: s,
	}
	putHandler := resourceGroupsService.PutHandler()
	getHandler := resourceGroupsService.PutHandler()

	// Assert that there were no errors and our resource groups service created and returned the new resource group
	if assert.NoError(t, putHandler(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), subscriptionID)
		assert.Contains(t, rec.Body.String(), resourceGroupName)
	}

	// Assert that there were no errors and our resource groups service returned the requested resource group
	if assert.NoError(t, getHandler(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), subscriptionID)
		assert.Contains(t, rec.Body.String(), resourceGroupName)
	}
}
