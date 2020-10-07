package lightdb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	expected := "{'a': 'b'}"

	store := NewStore()
	store.Put("/subscriptions/", "{'a': 'b'}")

	value := store.Get("/subscriptions/")
	assert.Equal(
		t,
		expected,
		value,
		"Got incorrect value from key. Expected %s, got: %s",
		expected,
		value,
	)

	store.Delete("/subscriptions/")

	value = store.Get("/subscriptions/")
	assert.Equal(t, value, "")
}

// TestGet shows a range of keys can be used
func TestGet(t *testing.T) {
	store := NewStore()
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
			val := store.Get(tt.key)
			assert.Equal(t, val, tt.expected, "Got unexpected value for key: %s, expected: %s, got: %s", tt.key, tt.expected, val)
		})
	}
}
