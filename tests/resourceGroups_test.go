package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-05-01/resources"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/goshlanguage/cerulean"
	"github.com/stretchr/testify/assert"
)

func TestPutResourceGroupsRoute(t *testing.T) {
	subscriptionID := "c27e7a81-b684-4fce-91d8-fed9e9bb534a"
	resourceGroupName := "ceruleanResourceGroup"

	server := cerulean.New("c27e7a81-b684-4fce-91d8-fed9e9bb534a")
	client := resources.NewGroupsClientWithBaseURI(server.GetBaseClientURI(), subscriptionID)
	client.Authorizer = autorest.NullAuthorizer{}

	resultPage, err := client.CreateOrUpdate(
		context.TODO(),
		resourceGroupName,
		resources.Group{
			Location: to.StringPtr("eastus2"),
		})
	if err != nil {
		panic(err)
	}

	fmt.Println(resultPage)
	_, err = resultPage.MarshalJSON()
	assert.NoErrorf(t, err, "Error raised when marshalling the client response: %s", err)

	// assert.Containsf(t, resultPage., "c27e7a81-b684-4fce-91d8-fed9e9bb534a", "Didn't find our created sub in the azure sdk subscriptions client List resultPage.")
}
