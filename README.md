<p align="center">
    <a href="https://aka.ms/free-account">
    <img src="https://raw.githubusercontent.com/ashleymcnamara/gophers/296b4d47f5313822b348e442837ca2d32a7704a3/Azure_Gophers.png" width="360"></a>
</p>

# Cerulean
Cerulean is a mock Azure API, designed for writing infrastructure tests before spending that dosh. This is great for folks who are using the `azure-sdk-for-go` especially. Here's a nifty example:

```go
package test

import (
    "fmt"
    "testing"

    "github.com/azure/azure-sdk-for-go"
    "github.com/goshlanguage/cerulean"
)

func TestSubscription(t *testing.T) {
    server := cerulean.New(":8080", "c27e7a81-b684-4fce-91d8-fed9e9bb534a")
}
```

Inspired by [moto](https://github.com/spulec/moto), Cerulean works by creating an http server that mimics Azure API endpoint responses. 



# Design
Cerulean emalates the Azure API. There are two good references for this:
 1. [azure-rest-api-specs](https://github.com/Azure/azure-rest-api-specs)
 1. [azure rest api docs](https://docs.microsoft.com/en-us/rest/api/azure/)

 If you decide to explore the API via curl or postman with the docs above, this article may be helpful to you:
 https://medium.com/@mauridb/calling-azure-rest-api-via-curl-eb10a06127

