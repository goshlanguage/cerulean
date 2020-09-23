package tests

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2016-10-01/keyvault"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/goshlanguage/cerulean"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

// TestKeyvaultCreateOrUpdateRoute creates a keyvault, then checks for its existence
func TestKeyvaultCreateOrUpdateRoute(t *testing.T) {
	server := cerulean.New()
	client := keyvault.NewVaultsClientWithBaseURI(server.GetBaseClientURI(), server.SubscriptionID)
	client.Authorizer = autorest.NullAuthorizer{}

	resourceGroup := "testRG"
	vaultName := "Vault101"
	tenantID, err := uuid.FromString(server.TenantID)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	client.CreateOrUpdate(
		context.Background(),
		resourceGroup,
		vaultName,
		keyvault.VaultCreateOrUpdateParameters{
			Location: to.StringPtr("eastus"),
			Properties: &keyvault.VaultProperties{
				TenantID: &tenantID,
				Sku: &keyvault.Sku{
					Family: to.StringPtr("A"),
					Name:   keyvault.Standard,
				},
				AccessPolicies: &[]keyvault.AccessPolicyEntry{},
			},
		},
	)

	assert.NotEqual(
		t,
		len(server.Inventory.Keyvaults),
		0,
	)
}
