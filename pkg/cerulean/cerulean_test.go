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
	cerulean := New("c27e7a81-b684-4fce-91d8-fed9e9bb534a")
	assert.Equal(t, (*cerulean.Subscriptions)[0].ID, "/subscriptions/c27e7a81-b684-4fce-91d8-fed9e9bb534a", "Received an invalid subscription id")

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
		"{\"value\":[{\"id\":\"/subscriptions/c27e7a81-b684-4fce-91d8-fed9e9bb534a\",\"authorizationSource\":\"RoleBased\",\"managedByTenants\":[],\"subscriptionId\":\"c27e7a81-b684-4fce-91d8-fed9e9bb534a\",\"tenantId\":\"b5549535-3215-4868-a289-f80095c9e718\",\"displayName\":\"Pay-As-You-Go\",\"state\":\"Enabled\",\"subscriptionPolicies\":{\"locationPlacementId\":\"Public_2014-09-01\",\"quotaId\":\"PayAsYouGo_2014-09-01\",\"spendingLimit\":\"Off\"}}],\"count\":{\"type\":\"\",\"value\":0}}",
		string(body))
}
