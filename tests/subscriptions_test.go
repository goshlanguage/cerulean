package tests

// TODO: Don't use /latest
import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-11-01/subscriptions"
	"github.com/Azure/go-autorest/autorest"
	"github.com/goshlanguage/cerulean"
	"github.com/stretchr/testify/assert"
)

// TestGetSubscriptionsRoute makes a request to the Cerulean API and expects that we get back our default subscriptionID
func TestGetSubscriptionsRoute(t *testing.T) {
	server := cerulean.New()
	client := subscriptions.NewClientWithBaseURI(server.GetBaseClientURI())
	client.Authorizer = autorest.NullAuthorizer{}

	var subs []subscriptions.Subscription
	for subList, err := client.List(context.Background()); subList.NotDone(); err = subList.Next() {
		if err != nil {
			t.Errorf("failed to get list of subs: %v", err)
			t.Fail()
		}
		subs = append(subs, subList.Values()...)
	}

	assert.Equal(
		t,
		*subs[0].SubscriptionID,
		server.SubscriptionID,
		"Didn't find our created sub in the azure sdk subscriptions.",
	)
}
