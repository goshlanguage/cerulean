package tests

// TODO: Don't use /latest
import (
	"context"
	"fmt"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-11-01/subscriptions"
	"github.com/Azure/go-autorest/autorest"
	"github.com/goshlanguage/cerulean/pkg/cerulean"
)

func TestGetSubscriptionsRoute(t *testing.T) {
	// Setup server
	//	assert.Equal(t, (*server.Subscriptions)[0].ID, "/subscriptions/c27e7a81-b684-4fce-91d8-fed9e9bb534a", "Received an invalid subscription id")
	server := cerulean.New("c27e7a81-b684-4fce-91d8-fed9e9bb534a")
	client := subscriptions.NewClientWithBaseURI(server.GetBaseClientURI())
	client.Authorizer = autorest.NullAuthorizer{}

	resultPage, err := client.List(context.TODO())
	if err != nil {
		panic(err)
	}
	// TODO lets render out this object and find out how and what to assert on
	fmt.Println(resultPage)
}
