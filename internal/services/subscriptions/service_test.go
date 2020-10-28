package subscriptions

import (
	"testing"

	"github.com/goshlanguage/cerulean/pkg/lightdb"
	"github.com/stretchr/testify/assert"
)

// TestAddAndGetSubscription creates a store and tests out the AddSubscription(s Subscription) helper to ensure it functions as expected
func TestAddSubscription(t *testing.T) {
	s := lightdb.NewStore()
	svc := NewService(s)

	newSub := NewSubscription()
	err := svc.AddSubscription(newSub)
	assert.NoError(t, err, "Adding subscription through service helper failed: %s", err)

	dbSubs := s.Get("subscriptions")
	assert.NotEmpty(t, dbSubs, "Expected a non empty response from lightdb.")
	assert.Contains(t, dbSubs, newSub.ID, "Expected to find newSub's ID in lightdb")
}

func TestGetSubscriptions(t *testing.T) {
	s := lightdb.NewStore()
	svc := NewService(s)

	subs, err := svc.GetSubscriptions()
	assert.NoError(t, err, "Tried to get subscriptions from service helper but received error: %s", err)
	assert.Equal(t, 1, len(subs), "Expected to initial subscription to exist after the service")

	for i := 0; i < 10; i++ {
		err := svc.AddSubscription(NewSubscription())
		assert.NoError(t, err, "Adding subscription through service helper failed: %s", err)
	}

	subs, err = svc.GetSubscriptions()
	assert.NoError(t, err, "Tried to get subscriptions from service helper but received error: %s", err)
	assert.Equal(t, 11, len(subs), "Expected to have 11 subscriptions after adding 10 of them to the inital subscription")
}
