package cerulean

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerInitilization(t *testing.T) {
	cerulean := New()
	assert.Equal(t, (*cerulean.Inventory.Subscriptions)[0].ID, fmt.Sprintf("/subscriptions/%s", cerulean.SubscriptionID), "Received an invalid subscription id")

	ts := httptest.NewServer(cerulean.Handlers["/subscriptions"])
	defer ts.Close()

	addr := fmt.Sprintf("%s/subscriptions", ts.URL)
	res, err := http.Get(addr)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t,
		fmt.Sprintf("{\"value\":[{\"id\":\"/subscriptions/%s\",\"authorizationSource\":\"RoleBased\",\"managedByTenants\":[],\"subscriptionId\":\"%s\",\"tenantId\":\"b5549535-3215-4868-a289-f80095c9e718\",\"displayName\":\"Pay-As-You-Go\",\"state\":\"Enabled\",\"subscriptionPolicies\":{\"locationPlacementId\":\"Public_2014-09-01\",\"quotaId\":\"PayAsYouGo_2014-09-01\",\"spendingLimit\":\"Off\"}}],\"count\":{\"type\":\"\",\"value\":0}}", cerulean.SubscriptionID, cerulean.SubscriptionID),
		string(body))
}

func TestListenAndServe(t *testing.T) {
	cerulean := New()
	cerulean.ListenAndServe()

	endpoint := fmt.Sprintf("%s/subscriptions", cerulean.GetBaseClientURI())
	resp, err := http.Get(endpoint)
	assert.NoErrorf(t, err, "Expected not to receive an error when requesting from the cerulean mock server, got: %s", err)
	assert.Equal(t, 200, resp.StatusCode, "Expected a 200 status code from the cerulean mock server")
}
