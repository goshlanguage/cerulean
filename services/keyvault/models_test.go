package keyvault

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewKeyvault(t *testing.T) {
	kv, _ := NewKeyvault("mykeyvault", "eastus")
	assert.Equal(t, kv.Name, "mykeyvault")
}
