The API response for /subscriptions look as follows (p.s. the uuids are fake, so don't try to hack the planet with these).
If you'd like to make your own API requests, you can follow this guide to get you started:

https://medium.com/@mauridb/calling-azure-rest-api-via-curl-eb10a06127

`GET`:

```sh
➜ curl -s -X GET -H "Authorization: Bearer $TOKEN" -H "Content-Type: application/json" "https://management.azure.com/subscriptions?api-version=2020-01-01" | jq .
{
  "value": [
    {
      "id": "/subscriptions/b5549535-3215-4868-a289-f80095c9e718",
      "authorizationSource": "RoleBased",
      "managedByTenants": [],
      "subscriptionId": "b5549535-3215-4868-a289-f80095c9e718",
      "tenantId": "b5549535-3215-4868-a289-f80095c9e718",
      "displayName": "Pay-As-You-Go",
      "state": "Enabled",
      "subscriptionPolicies": {
        "locationPlacementId": "Public_2014-09-01",
        "quotaId": "PayAsYouGo_2014-09-01",
        "spendingLimit": "Off"
      }
    }
  ],
  "count": {
    "type": "Total",
    "value": 1
  }
}
```

