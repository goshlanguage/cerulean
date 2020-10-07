package subscriptions

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetSubscriptionsHandler is the GET method handler for /subscriptions
func (svc *SubscriptionService) GetSubscriptionsHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		subsString := svc.Store.Get("subscriptions")

		// If something is stored in the db
		if subsString != "" {
			subs := []Subscription{}
			err := json.Unmarshal([]byte(subsString), &subs)
			if err != nil {
				panic(err)
			}

			response := SubscriptionResponse{
				Value: subs,
			}

			c.JSON(http.StatusOK, response)
		}
		return nil
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
