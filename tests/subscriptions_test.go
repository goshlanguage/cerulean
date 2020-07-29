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

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-11-01/subscriptions"
	"github.com/Azure/go-autorest/autorest"
	"github.com/goshlanguage/cerulean/pkg/cerulean"

	"github.com/stretchr/testify/assert"
)

func TestGetSubscriptionsRoute(t *testing.T) {
	// Setup SDK
	// ------
	// Setup nullauthorizer
	// Setup subscription client
	// Request subs
	// Validate that our sub exists

	// Setup server
	server := cerulean.New(":8080", "c27e7a81-b684-4fce-91d8-fed9e9bb534a")
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

	assert.Equal(t, "{\"value\":[{\"id\":\"/subscriptions/c27e7a81-b684-4fce-91d8-fed9e9bb534a\",\"subscriptionId\":\"c27e7a81-b684-4fce-91d8-fed9e9bb534a\",\"displayName\":\"Pay-As-You-Go\",\"state\":\"Enabled\",\"subscriptionPolicies\":{\"locationPlacementId\":\"Public_2014-09-01\",\"quotaId\":\"PayAsYouGo_2014-09-01\",\"spendingLimit\":\"Off\"}}]}", string(subscriptionResponse))

	client := subscriptions.NewClientWithBaseURI(ts.URL)
	client.Authorizer = autorest.NullAuthorizer{}

	resultPage, err := client.List(context.TODO())
	if err != nil {
		panic(err)
	}
	// TODO lets render out this object and find out how and what to assert on
	fmt.Println(resultPage)
}
