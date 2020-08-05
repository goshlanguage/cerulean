package cerulean

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServerInitilization(t *testing.T) {
	cerulean := New()
	assert.Regexp(
		t,
		regexp.MustCompile("[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{8}"),
		cerulean.BaseSubscriptionID,
		"Received an invalid subscription id: %s",
		cerulean.BaseSubscriptionID,
	)

	addr := fmt.Sprintf("%s/subscriptions", cerulean.GetBaseClientURI())
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
		fmt.Sprintf("{\"value\":[{\"id\":\"/subscriptions/%s\",\"authorizationSource\":\"RoleBased\",\"managedByTenants\":[],\"subscriptionId\":\"%s\",\"tenantId\":\"b5549535-3215-4868-a289-f80095c9e718\",\"displayName\":\"Pay-As-You-Go\",\"state\":\"Enabled\",\"subscriptionPolicies\":{\"locationPlacementId\":\"Public_2014-09-01\",\"quotaId\":\"PayAsYouGo_2014-09-01\",\"spendingLimit\":\"Off\"}}],\"count\":{\"type\":\"\",\"value\":0}}\n", cerulean.BaseSubscriptionID, cerulean.BaseSubscriptionID),
		string(body),
	)
}

func TestGetBaseClientURI(t *testing.T) {
	cerulean := New()
	addr := cerulean.GetBaseClientURI()
	assert.NotEmpty(t, addr, "Expected a populated address to exist")
	assert.Regexpf(
		t,
		regexp.MustCompile("http://127.0.0.1:[0-9]{2,5}"),
		addr,
		"Got inappropriate address for GetBaseClientURI(). Got: %s",
		addr,
	)
}

func TestListenAndServe(t *testing.T) {
	cerulean := New()

	endpoint := fmt.Sprintf("%s/subscriptions", cerulean.GetBaseClientURI())
	resp, err := http.Get(endpoint)

	assert.NoErrorf(t, err, "Expected not to receive an error when requesting from the cerulean mock server, got: %s", err)
	assert.Equal(t, 200, resp.StatusCode, "Expected a 200 status code from the cerulean mock server")
}
