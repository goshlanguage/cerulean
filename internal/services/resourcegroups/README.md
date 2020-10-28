# `/resourcegroups`

The API response for `/resourcegroups` is documented below.

If you'd like to make your own API requests, you can follow [this guide to get you started](https://medium.com/@mauridb/calling-azure-rest-api-via-curl-eb10a06127).

## `GET`

https://docs.microsoft.com/en-us/rest/api/resources/resourcegroups/get

```
GET https://management.azure.com/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}?api-version=2020-06-01

{
  "id": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/my-resource-group",
  "name": "my-resource-group",
  "location": "eastus",
  "properties": {
    "provisioningState": "Succeeded"
  }
}
```

## `PUT`

https://docs.microsoft.com/en-us/rest/api/resources/resourcegroups/createorupdate

```
PUT https://management.azure.com/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}?api-version=2020-06-01

{
  "id": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/my-resource-group",
  "name": "my-resource-group",
  "location": "eastus",
  "properties": {
    "provisioningState": "Succeeded"
  }
}
```
