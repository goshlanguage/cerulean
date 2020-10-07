package subscriptions

import (
	"testing"

	"github.com/goshlanguage/cerulean/pkg/lightdb"
	"github.com/stretchr/testify/assert"
)

// TestAddSubscription creates a store and tests out the AddSubscription(s Subscription) helper to ensure it functions as expected
func TestAddSubscription(t *testing.T) {
	s := lightdb.NewStore()
	svc := NewSubscriptionService(s)

	newSub := NewSubscription()
	err := svc.AddSubscription(newSub)
	assert.NoError(t, err, "Adding subscription through service helper failed: %s", err)

	dbSubs, err := s.Get("subscriptions")
	assert.NoError(t, err, "Tried to get subscriptions from lightdb but received error: %s", err)

	assert.NotEmpty(t, dbSubs, "Expected a non empty response from lightdb.")
	assert.Contains(t, dbSubs, newSub.ID, "Expected to find newSub's ID in lightdb")
}
