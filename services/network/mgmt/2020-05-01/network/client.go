package network

import (
	"github.com/Azure/go-autorest/autorest"
)

// BaseClient emulates the upstream network BaseClient
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// NewWithBaseURI is a factory for BaseClients
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent("cerulean"),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}
