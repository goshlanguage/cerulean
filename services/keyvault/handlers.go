package keyvault

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// PostCreateOrUpdate is the POST method handler for CreateOrUpdate.
// It takes in a body object with all of the relevant information to
// create or update a keyvault mock
// https://docs.microsoft.com/en-us/rest/api/keyvault/vaults/createorupdate#create-a-new-vault-or-update-an-existing-vault
func PostCreateOrUpdate(keyvaults *[]Keyvault) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		var createOrUpdateBody CreateOrUpdateBody
		err = json.Unmarshal(body, &createOrUpdateBody)
		if err != nil {
			// TODO real error handling
			// TODO response to request with proper errors
			panic(err)
		}

		uri := r.RequestURI
		uriStructure := strings.Split(uri, "/")
		name := uriStructure[len(uriStructure)-1]

		newKV, stub := NewKeyvault(name, createOrUpdateBody.Location)
		newKVS := append(*keyvaults, newKV)
		keyvaults = &newKVS

		b, err := json.Marshal(stub)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Charset", "UTF-8")
		w.Write(b)
	})
}
