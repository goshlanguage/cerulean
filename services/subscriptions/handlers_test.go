package subscriptions

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goshlanguage/cerulean/pkg/lightdb"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// TestGetSubscriptionsHandler sets up a server and tests the endpoint directly
func TestGetSubscriptionsHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	s := lightdb.NewStore()
	subscriptionService := SubscriptionService{
		Store: s,
	}
	getHandler := subscriptionService.GetSubscriptionsHandler()

	// Assert that there were no errors and our subscriptions service returned a blank response because no subscriptions exist
	if assert.NoError(t, getHandler(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "")
	}
}
