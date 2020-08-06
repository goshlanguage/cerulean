package tests

import (
	"context"
	"log"
	"testing"

	previewSubscriptions "github.com/Azure/azure-sdk-for-go/services/preview/subscription/mgmt/2019-10-01-preview/subscription"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-11-01/subscriptions"
	"github.com/Azure/go-autorest/autorest"
	"github.com/goshlanguage/cerulean"
	"github.com/stretchr/testify/assert"
)

func TestGetSubscriptionsRoute(t *testing.T) {
	server := cerulean.New()
	client := subscriptions.NewClientWithBaseURI(server.GetBaseClientURI())
	client.Authorizer = autorest.NullAuthorizer{}

	resultPage, err := client.List(context.TODO())
	if err != nil {
		panic(err)
	}

	_, err = resultPage.Values()[0].MarshalJSON()
	assert.NoErrorf(t, err, "Error raised when marshalling the client response: %s", err)

	assert.Containsf(t, *resultPage.Values()[0].ID, server.BaseSubscriptionID, "Didn't find our created sub in the Azure SDK subscriptions client List resultPage.")
	assert.Containsf(t, *resultPage.Values()[0].SubscriptionID, server.BaseSubscriptionID, "Didn't find our created sub in the Azure SDK subscriptions client List resultPage.")
}

func TestPostSubscriptionsRoute(t *testing.T) {
	server := cerulean.New()
	client := previewSubscriptions.NewClientWithBaseURI(server.GetBaseClientURI())
	client.Authorizer = autorest.NullAuthorizer{}

	subscriptionDisplayName := "displayName"
	subscriptionSKU := "sku"
	subscriptionBody := previewSubscriptions.ModernSubscriptionCreationParameters{DisplayName: &subscriptionDisplayName, SkuID: &subscriptionSKU}
	resultPage, err := client.CreateSubscription(context.TODO(), "billingAccount", "billingProfile", "invoiceSection", subscriptionBody)
	if err != nil {
		panic(err)
	}

	log.Printf("%+v", resultPage)

	/*
		_, err = resultPage.Values()[0].MarshalJSON()
		assert.NoErrorf(t, err, "Error raised when marshalling the client response: %s", err)

		assert.Containsf(t, *resultPage.Values()[0].ID, server.BaseSubscriptionID, "Didn't find our created sub in the Azure SDK subscriptions client List resultPage.")
		assert.Containsf(t, *resultPage.Values()[0].SubscriptionID, server.BaseSubscriptionID, "Didn't find our created sub in the Azure SDK subscriptions client List resultPage.")
	*/
}
