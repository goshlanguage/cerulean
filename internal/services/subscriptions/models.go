package subscriptions

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

var sampleJSON = `
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
`

// Response models the subscription response from the API
type Response struct {
	Value []Subscription `json:"value"`
	Count struct {
		Type  string `json:"type"`
		Value int    `json:"value"`
	} `json:"count"`
}

// Subscription is the object we store in our Inventory grab bag to model a subscription
type Subscription struct {
	ID                   string        `json:"id"`
	AuthorizationSource  string        `json:"authorizationSource"`
	ManagedByTenants     []interface{} `json:"managedByTenants"`
	SubscriptionID       string        `json:"subscriptionId"`
	TenantID             string        `json:"tenantId"`
	DisplayName          string        `json:"displayName"`
	State                string        `json:"state"`
	SubscriptionPolicies struct {
		LocationPlacementID string `json:"locationPlacementId"`
		QuotaID             string `json:"quotaId"`
		SpendingLimit       string `json:"spendingLimit"`
	} `json:"subscriptionPolicies"`
}

// NewResponseStub takes a string ID and returns a basic SubscriptionResponse object
// TODO: Support passing multiple subscription IDs
func NewResponseStub(subscriptionID string) Response {
	var response Response
	json.Unmarshal([]byte(sampleJSON), &response)

	response.Value[0].ID = fmt.Sprintf("/subscriptions/%s", subscriptionID)
	response.Value[0].SubscriptionID = fmt.Sprintf("%s", subscriptionID)

	return response
}

// NewSubscription takes a string ID and returns a basic Subscription object
// It does this by indirectly using the NewSubscriptionResponse factory for the sake of loading in a JSON stub,
//   and replacing the SubscriptionID in said stub to constructr our Subscription.
func NewSubscription() Subscription {
	// Generate a UUID for the SubscriptionID
	id, err := uuid.NewUUID()
	// TODO better error handling
	if err != nil {
		panic(err)
	}

	return NewResponseStub(id.String()).Value[0]
}
