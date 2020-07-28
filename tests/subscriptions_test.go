package tests

// TODO: Don't use /latest
import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	azureSubscriptions "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-11-01/subscriptions"
	"github.com/Azure/go-autorest/autorest"
	"github.com/goshlanguage/cerulean/pkg/server"
	"github.com/goshlanguage/cerulean/pkg/services/subscriptions"

	"github.com/stretchr/testify/assert"
)

func TestApiCallToCerulean(t *testing.T) {
	// Setup SDK
	// ------
	// Setup nullauthorizer
	// Setup subscription client
	// Request subs
	// Validate that our sub exists

	// TODO: Make helper that generates a sub
	// Setup server
	server := server.GetServer(":8080")
	// Add subscription to inventory
	*server.Subscriptions = append(*server.Subscriptions, subscriptions.Subscription{
		ID:             "/subscriptions/c27e7a81-b684-4fce-91d8-fed9e9bb534a",
		SubscriptionID: "c27e7a81-b684-4fce-91d8-fed9e9bb534a",
		DisplayName:    "mysub",
		State:          "Enabled",
	})

	assert.Equal(t, (*server.Subscriptions)[0].ID, "/subscriptions/c27e7a81-b684-4fce-91d8-fed9e9bb534a", "Received an invalid subscription id")

	ts := httptest.NewServer(server.Handlers["/subscriptions"])
	defer ts.Close()

	addr := fmt.Sprintf("%s/subscriptions", ts.URL)
	res, err := http.Get(addr)
	if err != nil {
		log.Fatal(err)
	}

	subscriptionResponse, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, "[{\"id\":\"/subscriptions/c27e7a81-b684-4fce-91d8-fed9e9bb534a\",\"subscriptionId\":\"c27e7a81-b684-4fce-91d8-fed9e9bb534a\",\"displayName\":\"mysub\",\"state\":\"Enabled\",\"subscriptionPolicies\":{\"locationPlacementId\":\"\",\"quotaId\":\"\",\"spendingLimit\":\"\"}}]", string(subscriptionResponse))

	client := azureSubscriptions.NewClientWithBaseURI(ts.URL)
	client.Authorizer = autorest.NullAuthorizer{}

	resultPage, err := client.List(context.TODO())
	if err != nil {
		panic(err)
	}

	fmt.Println(resultPage)

	assert.True(t, false)
}
