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
	server := New()
	assert.Equal(
		t,
		(*server.Inventory.Subscriptions)[0].ID,
		fmt.Sprintf("/subscriptions/%s", server.SubscriptionID),
		"Received an invalid subscription id",
	)

	assert.NotNil(
		t,
		server.Handlers["/subscriptions/"],
		"/subscriptions subscription handler not populated.",
	)

	ts := httptest.NewServer(server.Handlers["/subscriptions/"])
	defer ts.Close()

	addr := fmt.Sprintf("%s/subscriptions/", ts.URL)
	res, err := http.Get(addr)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	assert.Containsf(
		t,
		string(body),
		fmt.Sprintf(`{"value":[{"id":"/subscriptions/%s"`, server.SubscriptionID),
		fmt.Sprintf("Subscription response didn't return default subscription. Got: %s from %s", string(body), addr),
	)
}

func TestListenAndServe(t *testing.T) {
	cerulean := New()
	cerulean.ListenAndServe()

	endpoint := fmt.Sprintf("%s/subscriptions", cerulean.GetBaseClientURI())
	resp, err := http.Get(endpoint)
	assert.NoErrorf(t, err, "Expected not to receive an error when requesting from the cerulean mock server, got: %s", err)
	assert.Equal(t, 200, resp.StatusCode, "Expected a 200 status code from the cerulean mock server")
}
