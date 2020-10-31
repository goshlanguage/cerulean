package subscriptions

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetHandler is the GET method handler for /subscriptions
func (svc *Service) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		subsString := svc.Store.Get("subscriptions")

		// If something is stored in the db
		if subsString != "" {
			subs := []Subscription{}
			err := json.Unmarshal([]byte(subsString), &subs)
			if err != nil {
				panic(err)
			}

			response := Response{
				Value: subs,
			}

			c.JSON(http.StatusOK, response)
		}
		return nil
	}
}

// PostHandler is the POST method handler for /subscriptions
// POST https://management.azure.com/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Subscription/createSubscription?api-version=2018-11-01-preview
// 202 Accepted
// Response: {
// 	"subscriptionLink": "/subscriptions/d0d6ee57-6530-4fca-93a6-b755a070be35"
// }
func PostHandler(pattern string, subs *[]Subscription) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
