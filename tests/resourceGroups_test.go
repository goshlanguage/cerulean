package tests

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/goshlanguage/cerulean"
	"github.com/stretchr/testify/assert"
)

func TestPutResourceGroupsRoute(t *testing.T) {
	server := cerulean.New()
	client := resources.NewGroupsClientWithBaseURI(server.GetBaseClientURI(), server.BaseSubscriptionID)
	client.Authorizer = autorest.NullAuthorizer{}

	resourceGroupName := "testResourceGroup"
	resourceGroupLocation := "westus"

	resourceGroup, err := client.CreateOrUpdate(
		context.TODO(),
		resourceGroupName,
		resources.Group{
			Location: to.StringPtr(resourceGroupLocation)})
	if err != nil {
		panic(err)
	}

	assert.Equal(t, resourceGroup.Name, resourceGroupName, "Created resource group name does not match expected value")
	assert.Equal(t, resourceGroup.Location, resourceGroupLocation, "Created resource group location does not match expected value")
}
