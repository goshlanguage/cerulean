package keyvault

import (
	"encoding/json"
	"io/ioutil"

	"github.com/google/uuid"
)

// CheckNameAvailabilityResponse is the data model for the Azure API response to
// subscriptions/$SUBSCRIPTION_ID/providers/Microsoft.KeyVault/checkNameAvailability
type CheckNameAvailabilityResponse struct {
	NameAvailable bool `json:"nameAvailable"`
}

// CreateOrUpdateResponse allows us to marshal our stub into a response
type CreateOrUpdateResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Location string `json:"location"`
	Tags     struct {
	} `json:"tags"`
	Properties Properties `json:"properties"`
}

// CreateOrUpdateBody is used to unmarshal the expected body object for this endpoint
type CreateOrUpdateBody struct {
	Location   string     `json:"location"`
	Properties Properties `json:"properties"`
}

// Properties is broken out for embedding into other projects
type Properties struct {
	Sku struct {
		Family string `json:"family"`
		Name   string `json:"name"`
	} `json:"sku"`
	TenantID                     string         `json:"tenantId"`
	AccessPolicies               []AccessPolicy `json:"accessPolicies"`
	EnabledForDeployment         bool           `json:"enabledForDeployment"`
	EnabledForDiskEncryption     bool           `json:"enabledForDiskEncryption"`
	EnabledForTemplateDeployment bool           `json:"enabledForTemplateDeployment"`
	VaultURI                     string         `json:"vaultUri"`
}

// AccessPolicy is broken out for embedding in other objects
type AccessPolicy struct {
	TenantID    string `json:"tenantId"`
	ObjectID    string `json:"objectId"`
	Permissions struct {
		Keys         []string `json:"keys"`
		Secrets      []string `json:"secrets"`
		Certificates []string `json:"certificates"`
	} `json:"permissions"`
}

// Keyvault is our inventory model for a Keyvault
type Keyvault struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Location string `json:"location"`
	Tags     struct {
	} `json:"tags"`
	Properties Properties `json:"properties"`
}

// NewKeyvault is a helper that constructs a Keyvault object, and returns a Response object
func NewKeyvault(name string, location string) (Keyvault, CreateOrUpdateResponse) {
	id, err := uuid.NewUUID()
	if err != nil {
		// TODO real error handling
		panic(err)
	}

	file, err := ioutil.ReadFile("stubs/createorupdate-201.json")
	if err != nil {
		panic(err)
	}

	stub := CreateOrUpdateResponse{}
	err = json.Unmarshal([]byte(file), &stub)

	stub.Name = name
	stub.ID = id.String()
	stub.Location = location

	return Keyvault{
		ID:         id.String(),
		Name:       name,
		Type:       "Microsoft.KeyVault/vaults",
		Location:   location,
		Properties: stub.Properties,
	}, stub
}
