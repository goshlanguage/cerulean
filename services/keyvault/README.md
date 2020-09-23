# `/keyvault`

API responses for the `/subscriptions/$SUBSCRIPTION_ID/providers/Microsoft.KeyVault/` endpoints are documented below.


## Check Name Availability - POST

https://docs.microsoft.com/en-us/rest/api/keyvault/vaults/checknameavailability

```sh
➜ curl -s -X POST -d '{"name":"cerulean", "type":"Microsoft.KeyVault/vaults"}' -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" "https://management.azure.com/subscriptions/$SUBSCRIPTION_ID/providers/Microsoft.KeyVault/checkNameAvailability?api-version=2019-09-01" | jq .
{
  "nameAvailable": true
}
```

## Create or Update - PUT

```
PUT https://management.azure.com/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-resource-group/providers/Microsoft.KeyVault/vaults/sample-vault?api-version=2019-09-01
```

https://docs.microsoft.com/en-us/rest/api/keyvault/vaults/createorupdate#vaultproperties

```json
{
  "id": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/sample-resource-group/providers/Microsoft.KeyVault/vaults/sample-vault",
  "name": "sample-vault",
  "type": "Microsoft.KeyVault/vaults",
  "location": "westus",
  "tags": {},
  "properties": {
    "sku": {
      "family": "A",
      "name": "standard"
    },
    "tenantId": "00000000-0000-0000-0000-000000000000",
    "accessPolicies": [
      {
        "tenantId": "00000000-0000-0000-0000-000000000000",
        "objectId": "00000000-0000-0000-0000-000000000000",
        "permissions": {
          "keys": [
            "encrypt",
            "decrypt",
            "wrapKey",
            "unwrapKey",
            "sign",
            "verify",
            "get",
            "list",
            "create",
            "update",
            "import",
            "delete",
            "backup",
            "restore",
            "recover",
            "purge"
          ],
          "secrets": [
            "get",
            "list",
            "set",
            "delete",
            "backup",
            "restore",
            "recover",
            "purge"
          ],
          "certificates": [
            "get",
            "list",
            "delete",
            "create",
            "import",
            "update",
            "managecontacts",
            "getissuers",
            "listissuers",
            "setissuers",
            "deleteissuers",
            "manageissuers",
            "recover",
            "purge"
          ]
        }
      }
    ],
    "enabledForDeployment": true,
    "enabledForDiskEncryption": true,
    "enabledForTemplateDeployment": true,
    "vaultUri": "https://sample-vault.vault.azure.net"
  }
}
```

## Reference

|Comment|Link|
|-|-|
|Keyvault documentation|https://docs.microsoft.com/en-us/rest/api/keyvault/|
|JSON-to-Golang|https://mholt.github.io/json-to-go/|
|Check name availability|https://docs.microsoft.com/en-us/rest/api/keyvault/vaults/checknameavailability|
