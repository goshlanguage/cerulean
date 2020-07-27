package main

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2017-09-01/network"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

func main() {
	vnetClient := GetVNetClient("")

	// call the VirtualNetworks CreateOrUpdate API
	vnetClient.CreateOrUpdate(context.Background(),
		"<resourceGroupName>",
		"<vnetName>",
		network.VirtualNetwork{
			Location: to.StringPtr("<azureRegion>"),
			VirtualNetworkPropertiesFormat: &network.VirtualNetworkPropertiesFormat{
				AddressSpace: &network.AddressSpace{
					AddressPrefixes: &[]string{"10.0.0.0/8"},
				},
				Subnets: &[]network.Subnet{
					{
						Name: to.StringPtr("<subnet1Name>"),
						SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
							AddressPrefix: to.StringPtr("10.0.0.0/16"),
						},
					},
					{
						Name: to.StringPtr("<subnet2Name>"),
						SubnetPropertiesFormat: &network.SubnetPropertiesFormat{
							AddressPrefix: to.StringPtr("10.1.0.0/16"),
						},
					},
				},
			},
		})
}

// GetVNetClient takes a subscriptionID and returns a virtualnetworkclient for that subscription
func GetVNetClient(subscriptionID string) network.VirtualNetworksClient {
	// create a VirtualNetworks client
	vnetClient := network.NewVirtualNetworksClient(subscriptionID)

	// create an authorizer from env vars or Azure Managed Service Idenity
	vnetClient.Authorizer = autorest.NullAuthorizer{}
	return vnetClient
}
