package resourceGroups

import (
	"encoding/json"
	"net/http"
)

// PutResourceGroupsHandler is the PUT method handler for /subscriptions/{subscription-id}/resourceGroups
func PutResourceGroupsHandler(subscriptionID string, resourceGroupName string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := json.Marshal(NewResourceGroupsResponse(subscriptionID, resourceGroupName))
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Charset", "UTF-8")
		w.Write(b)
	})
}
