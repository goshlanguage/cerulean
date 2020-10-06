package subscriptions

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetSubscriptionsHandler is the GET method handler for /subscriptions
func (svc *SubscriptionService) GetSubscriptionsHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		response := SubscriptionResponse{
			Value: svc.Subscriptions,
		}

		return c.JSON(http.StatusOK, response)
	}
}

// PostSubscriptionsHandler is the POST method handler for /subscriptions
// POST https://management.azure.com/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Subscription/createSubscription?api-version=2018-11-01-preview
// 202 Accepted
// Response: {
// 	"subscriptionLink": "/subscriptions/d0d6ee57-6530-4fca-93a6-b755a070be35"
// }
// TODO: Figure out how we're going to handle the api-versioning query param
func (svc *SubscriptionService) PostSubscriptionsHandler() echo.HandlerFunc {
	type response struct {
		SubscriptionLink string `json:"subscriptionLink"`
	}

	return func(c echo.Context) error {
		newSubscription := NewSubscription()
		svc.Subscriptions = append(svc.Subscriptions, newSubscription)
		// The HTTP "Location" header has to be set to a non-empty dummy value
		c.Response().Header().Set(echo.HeaderLocation, "https://cerulean")
		c.Response().WriteHeader(http.StatusCreated)
		return json.NewEncoder(c.Response()).Encode(response{newSubscription.ID})
	}
}
