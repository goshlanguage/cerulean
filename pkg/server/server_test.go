package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/goshlanguage/cerulean/pkg/services/subscriptions"
	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	server := GetServer(":8080")
	server.Subscriptions = append(server.Subscriptions, &subscriptions.Subscription{
		ID:             "/subscriptions/c27e7a81-b684-4fce-91d8-fed9e9bb534a",
		SubscriptionID: "c27e7a81-b684-4fce-91d8-fed9e9bb534a",
		DisplayName:    "mysub",
		State:          "Enabled",
	})

	ts := httptest.NewServer(server.Handlers["/subscriptions/"])
	defer ts.Close()

	addr := fmt.Sprintf("%s/subscriptions/", ts.URL)
	t.Errorf("Addr: %s", addr)
	res, err := http.Get(addr)
	if err != nil {
		log.Fatal(err)
	}

	greeting, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, "Hello, world!\n", string(greeting))
}
