package lightdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	expected := "{'a': 'b'}"

	s := NewStore()
	s.Put("/subscriptions/", "{'a': 'b'}")

	value, err := s.Get("/subscriptions/")
	assert.NoError(t, err, "Didn't expect error when getting previously stored key, got: %s", err)
	assert.Equal(
		t,
		expected,
		value,
		"Got incorrect value from key. Expected %s, got: %s",
		expected,
		value,
	)

	err = s.Delete("/subscriptions/")
	assert.NoError(
		t,
		err,
		"Expected no errors from Delete. Error: %s",
		err,
	)

	_, err = s.Get("/subscriptions/")
	assert.EqualError(
		t,
		err,
		"Key does not exist",
	)
}
