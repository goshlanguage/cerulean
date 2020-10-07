package subscriptions

import (
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
func PostSubscriptionsHandler(pattern string, subs *[]Subscription) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
