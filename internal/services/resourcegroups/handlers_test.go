package resourcegroups

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TestPutHandler sets up a server and tests the endpoint directly
func TestPutHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/subscriptions/:subscriptionId/resourcegroups/:resourceGroupName", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	resourceGroupsService := Service{}
	getHandler := resourceGroupsService.PutHandler()

	// Assertions
	if assert.NoError(t, getHandler(ctx)) {
		// assert.Equal(t, http.StatusCreated, rec.Code)
	}
}
