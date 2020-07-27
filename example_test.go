package main

import (
	"context"
	"testing"

	"github.com/goshlanguage/cerulean/services/network/mgmt/2020-05-01/network"
	"github.com/stretchr/testify/assert"
)

func TestGetVNetClient(t *testing.T) {
	vnetClient := GetVNetClient("1234abcd")
	assert.NotNil(t, vnetClient.Authorizer, "Need a populated authorizer to continue")

	result, err := vnetClient.List(context.TODO(), "fakeRG")
	assert.Equal(t, result.Response().StatusCode, 404, "Expected not to find fake subscriptions and resource groups")
	assert.Error(t, err, "Expected autorest not to find our fake subscription when listing networks")
}

func TestFakeVNetClient(t *testing.T) {
	subscriptionID := "superfake"
	vnetClient := network.FakeNewVirtualNetworksClient(subscriptionID)
	result, err := vnetClient.List(context.TODO(), "fakeRG")
	assert.Equal(t, result.Response().StatusCode, 200, "Failed to receive a successful response code from mock")
	assert.NoError(t, err, "Expected no error when listing networks")
}
