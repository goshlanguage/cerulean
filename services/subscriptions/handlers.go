package subscriptions

import (
	"encoding/json"
	"net/http"
)

/*
GET:
➜ curl -s -X GET -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" "https://management.azure.com/subscriptions?api-version=2020-01-01" | jq .
{
  "value": [
    {
      "id": "/subscriptions/b5549535-3215-4868-a289-f80095c9e718",
      "authorizationSource": "RoleBased",
      "managedByTenants": [],
      "subscriptionId": "b5549535-3215-4868-a289-f80095c9e718",
      "tenantId": "b5549535-3215-4868-a289-f80095c9e718",
      "displayName": "Pay-As-You-Go",
      "state": "Enabled",
      "subscriptionPolicies": {
        "locationPlacementId": "Public_2014-09-01",
        "quotaId": "PayAsYouGo_2014-09-01",
        "spendingLimit": "Off"
      }
    }
  ],
  "count": {
    "type": "Total",
    "value": 1
  }
}
*/

// GetSubscriptionsHandler is the GET method handler for /subscriptions
func GetSubscriptionsHandler(subs *[]Subscription) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := SubscriptionResponse{
			Value: *subs,
		}

		b, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Charset", "UTF-8")
		// json.NewEncoder(w).Encode(b)
		w.Write(b)
	})
}

// PostSubscriptionsHandler is the POST method handler for /subscriptions
// POST https://management.azure.com/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}/providers/Microsoft.Subscription/createSubscription?api-version=2018-11-01-preview
// 202 Accepted
// Response: {
// 	"subscriptionLink": "/subscriptions/d0d6ee57-6530-4fca-93a6-b755a070be35"
// }
func PostSubscriptionsHandler(subs *[]Subscription) http.Handler {
	return http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
	})
}
