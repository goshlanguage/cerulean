package subscriptions

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TestGetSubscriptionsHandler sets up a server and tests the endpoint directly
func TestGetSubscriptionsHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	subscriptionService := SubscriptionService{}
	getHandler := subscriptionService.GetSubscriptionsHandler()

	// Assertions
	if assert.NoError(t, getHandler(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "{\"value\":null,\"count\":{\"type\":\"\",\"value\":0}}")
	}
}
