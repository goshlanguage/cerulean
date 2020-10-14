package resourcegroups

import (
	"encoding/json"
	"fmt"
)

var sampleJSON = `
{
  "id": "/subscriptions/cc6b141e-6afc-4786-9bf6-e3b9a5601460/resourceGroups/DefaultResourceGroup-EUS",
  "name": "DefaultResourceGroup-EUS",
  "location": "eastus",
  "properties": {
    "provisioningState": "Succeeded"
  }
}
`

// ResourceGroup represents a resource group and also models the response from the Azure API when creating or updating a resource group
type ResourceGroup struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Location   string `json:"location"`
	Properties struct {
		ProvisioningState string `json:"provisioningState"`
	} `json:"properties"`
}

// NewResourceGroupsResponse takes a string ID and returns a basic ResourceGroupsResponse object
func NewResourceGroupsResponse(subscriptionID string, resourceGroupName string) ResourceGroup {
	var response ResourceGroup
	json.Unmarshal([]byte(sampleJSON), &response)

	response.ID = fmt.Sprintf("/subscriptions/%s/resourceGroups/%s", subscriptionID, resourceGroupName)
	response.Name = resourceGroupName

	return response
}
