<p align="center">
    <a href="https://aka.ms/free-account">
    <img src="https://raw.githubusercontent.com/ashleymcnamara/gophers/296b4d47f5313822b348e442837ca2d32a7704a3/Azure_Gophers.png" width="360"></a>
</p>

# Cerulean

[![WIP](https://img.shields.io/badge/alpha-unstable-yellow)]() [![Go Report Card](https://goreportcard.com/badge/github.com/goshlanguage/cerulean)](https://goreportcard.com/report/github.com/goshlanguage/cerulean) [![Man Hours](https://img.shields.io/endpoint?url=https%3A%2F%2Fmh.jessemillar.com%2Fhours%3Frepo%3Dhttps%3A%2F%2Fgithub.com%2Fgoshlanguage%2Fcerulean.git)](https://jessemillar.com/r/man-hours)

> This project is [currently unmaintained](POSTMORTEM.md).

Cerulean is a mock Azure API, designed for writing infrastructure tests so you can run tests without spending that `dosh`. This is great for folks who are using the `azure-sdk-for-go` especially. Here's a nifty example:

```go
package test

import (
    "fmt"
    "testing"

    "github.com/azure/azure-sdk-for-go/services/resources/mgmt/2019-11-01/subscriptions"
    "github.com/goshlanguage/cerulean"
)

func TestSubscription(t *testing.T) {
    server := cerulean.New("c27e7a81-b684-4fce-91d8-fed9e9bb534a")
    client := subscriptions.NewClientWithBaseURI(server.GetBaseClientURI())
    client.Authorizer = autorest.NullAuthorizer{}

    resultPage, err := client.List(context.TODO())
    if err != nil {
	    panic(err)
    }
    // do stuff with your mocked client
}
```

This project is artisanally crafted, laptop to market software. Please open an issue if you find a bug or have a request. PRs are welcome.

# Design

Cerulean works by creating an http server that mimics the Azure API rest endpoints and responses. You can then redirect your Azure SDK clients to use the address of this mock server, and your client will then talk to this mock. This project is inspired by [moto](https://github.com/spulec/moto).

Cerulean emulates the Azure API. There are two good references for the API to use in development:
 1. [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs)
 1. [azure rest api docs](https://docs.microsoft.com/en-us/rest/api/azure/)

If you decide to explore the API via curl or postman with the docs above, this article may be helpful to you:
https://medium.com/@mauridb/calling-azure-rest-api-via-curl-eb10a06127

If you prefer [swagger](https://editor.swagger.io/), then you can load in the specs you want to reference by passing it's URL, such as:

- [subscriptions](https://raw.githubusercontent.com/Azure/azure-rest-api-specs/master/specification/subscription/resource-manager/Microsoft.Subscription/stable/2020-01-01/subscriptions.json)

