/*
The API response for /subscriptions look as follows (p.s. the uuids are fake, so good luck authenticating HACKERS)

GET:
{
  "value": [
    {
      "id": "/subscriptions/8c428685-0136-4981-9d65-8e1b1ef5f055",
      "subscriptionId": "df9b4673-18bf-4462-b824-2aee4c994c9f",
      "displayName": "Pay-As-You-Go",
      "state": "Enabled",
      "subscriptionPolicies": {
        "locationPlacementId": "Public_2014-09-01",
        "quotaId": "PayAsYouGo_2014-09-01",
        "spendingLimit": "Off"
      }
    }
  ]
}
*/

package subscriptions

// SubscriptionResponse models the subscription response from the API
type SubscriptionResponse struct {
	Value []struct {
		ID                   string `json:"id"`
		SubscriptionID       string `json:"subscriptionId"`
		DisplayName          string `json:"displayName"`
		State                string `json:"state"`
		SubscriptionPolicies struct {
			LocationPlacementID string `json:"locationPlacementId"`
			QuotaID             string `json:"quotaId"`
			SpendingLimit       string `json:"spendingLimit"`
		} `json:"subscriptionPolicies"`
	} `json:"value"`
}
