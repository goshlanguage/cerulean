package subscriptions

import (
	"fmt"
)

// SubscriptionResponse models the subscription response from the API
type SubscriptionResponse struct {
	Value []Subscription `json:"value"`
}

// Subscription is the object we store in our Inventory grab bag to model a subscription
type Subscription struct {
	ID                   string               `json:"id"`
	SubscriptionID       string               `json:"subscriptionId"`
	DisplayName          string               `json:"displayName"`
	State                string               `json:"state"`
	SubscriptionPolicies SubscriptionPolicies `json:"subscriptionPolicies"`
}

// SubscriptionPolicies is a model for subscriptions, and allows us to easily modify Subscription responses
type SubscriptionPolicies struct {
	LocationPlacementID string `json:"locationPlacementId"`
	QuotaID             string `json:"quotaId"`
	SpendingLimit       string `json:"spendingLimit"`
}

// NewSubscription takes a string ID and returns a basic Subscription object
func NewSubscription(subscriptionID string) Subscription {
	return Subscription{
		ID:             fmt.Sprintf("/subscriptions/%s", subscriptionID),
		SubscriptionID: fmt.Sprintf("%s", subscriptionID),
		DisplayName:    "Pay-As-You-Go",
		State:          "Enabled",
		SubscriptionPolicies: SubscriptionPolicies{
			LocationPlacementID: "Public_2014-09-01",
			QuotaID:             "PayAsYouGo_2014-09-01",
			SpendingLimit:       "Off",
		},
	}
}
