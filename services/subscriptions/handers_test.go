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
	req := httptest.NewRequest(http.MethodGet, "/subscriptions", nil)
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

// TestPostSubscriptionsHandler sets up a server and tests the endpoint directly
func TestPostSubscriptionsHandler(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/providers/Microsoft.Billing/billingAccounts/testAccount/billingProfiles/testProfile/invoiceSections/testSection/providers/Microsoft.Subscription/createSubscription?api-version=2018-11-01-preview", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	subscriptionService := SubscriptionService{}
	getHandler := subscriptionService.PostSubscriptionsHandler()

	// Assertions
	if assert.NoError(t, getHandler(ctx)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		// TODO: Maybe test that a UUID is returned
		assert.Contains(t, rec.Body.String(), "{\"subscriptionLink\":\"/subscriptions/")
	}
}
