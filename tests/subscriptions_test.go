package tests

// TODO: Don't use /latest
import (
	"context"
	"fmt"
	"testing"

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
	// TODO lets render out this object and find out how and what to assert on
	fmt.Println(resultPage)
	_, err = resultPage.Values()[0].MarshalJSON()
	assert.NoErrorf(t, err, "Error raised when marshalling the client response: %s", err)
	// assert.Containsf(t, json, cerulean.SubscriptionID, "Didn't find our created sub in the azure sdk subscriptions client List resultPage.")
}
