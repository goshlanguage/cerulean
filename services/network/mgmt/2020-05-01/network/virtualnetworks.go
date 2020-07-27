package network

const (
	DefaultBaseURI = "localhost"
)

type FakeVirtualNetworksClient struct {
	BaseClient
}

// NewVirtualNetworksClient mimics the upstream network factory function
func (client FakeVirtualNetworksClient) NewVirtualNetworksClient(subscriptionID string) FakeVirtualNetworksClient {
	return NewFakeVirtualNetworksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewFakeVirtualNetworksClientWithBaseURI creates an instance of the FakeVirtualNetworksClient client using a custom endpoint.
// We can use this to emulate interacting with an Azure cloud, by using a non-standard base URI.
func NewFakeVirtualNetworksClientWithBaseURI(baseURI string, subscriptionID string) FakeVirtualNetworksClient {
	return FakeVirtualNetworksClient{NewWithBaseURI(baseURI, subscriptionID)}
}
