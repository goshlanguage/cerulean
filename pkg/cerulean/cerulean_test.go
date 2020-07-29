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

	assert.Equal(t, "{\"value\":[{\"id\":\"/subscriptions/c27e7a81-b684-4fce-91d8-fed9e9bb534a\",\"subscriptionId\":\"c27e7a81-b684-4fce-91d8-fed9e9bb534a\",\"displayName\":\"Pay-As-You-Go\",\"state\":\"Enabled\",\"subscriptionPolicies\":{\"locationPlacementId\":\"Public_2014-09-01\",\"quotaId\":\"PayAsYouGo_2014-09-01\",\"spendingLimit\":\"Off\"}}]}", string(body))
}

func TestListenAndServe(t *testing.T) {
	cerulean := New("c27e7a81-b684-4fce-91d8-fed9e9bb534a")
	cerulean.ListenAndServe()

	endpoint := fmt.Sprintf("%s/subscriptions", cerulean.GetBaseClientURI())
	resp, err := http.Get(endpoint)
	assert.NoErrorf(t, err, "Expected not to receive an error when requesting from the cerulean mock server, got: %s", err)
	assert.Equal(t, 200, resp.StatusCode, "Expected a 200 status code from the cerulean mock server")
}
