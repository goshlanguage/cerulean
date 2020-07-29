package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// UnmarshalTemplate exists to remove a few lines of boilerplace from struct generation funcions in the subscriptions package
func UnmarshalTemplate(filename string) []byte {
	path, err := filepath.Abs(filename)
	if err != nil {
		return nil
	}

	// TODO: Maybe don't panic in this function; more graceful handling with a human-readable error?
	jsonFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer jsonFile.Close()

	contents, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}

	return contents
}
