package lightdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	expected := "{'a': 'b'}"

	s := NewStore()
	s.Put("/subscriptions/", "{'a': 'b'}")

	v, err := s.Get("/subscriptions/")
	assert.NoError(t, err, "Didn't expect error when getting previously stored key, got: %s", err)
	assert.Equal(
		t,
		expected,
		v,
		"Got incorrect value from key. Expected %s, got: %s",
		expected,
		v,
	)

	err = s.Delete("/subscriptions/")
	assert.NoError(
		t,
		err,
		"Expected no errors from Delete. Error: %s",
		err,
	)

	v, err = s.Get("/subscriptions/")
	assert.Equal(t, v, "")
	assert.NoError(t, err, "Didn't expect error from getting empty key")
}

// TestGet shows a range of keys can be used
func TestGet(t *testing.T) {
	s := NewStore()
	gettests := []struct {
		key      string
		expected string
		err      error
	}{
		{"invalid", "", nil},
		{"1234356", "", nil},
		{"-+#$)*(@!#", "", nil},
		{"🤔", "", nil},
		{"≈ç√∫˜µ≤åß∂ƒ©", "", nil},
	}
	for _, tt := range gettests {
		t.Run(tt.key, func(t *testing.T) {
			val, err := s.Get(tt.key)
			if err != nil {
				assert.Error(t, err, "Expected error: %s but got: %s", err, tt.err)
			} else {
				assert.NoError(t, err, "Didn't expect error from fetching key: %s", tt.key)
			}
			assert.Equal(t, val, tt.expected, "Got unexpected value for key: %s, expected: %s, got: %s", tt.key, tt.expected, val)
		})
	}
}
