package lightdb

import "fmt"

type Store struct {
	Data map[string]string
}

// NewStore is a factory for our store
func NewStore() *Store {
	return &Store{
		Data: make(map[string]string),
	}
}

// Put adds a new key value pair to the store or updates the key if present
func (s *Store) Put(key string, value string) error {
	s.Data[key] = value
	return nil
}

// Get returns the value for the key if it exists
func (s *Store) Get(key string) (string, error) {
	if val, ok := s.Data[key]; ok {
		return val, nil
	}
	return "", fmt.Errorf("Key does not exist")
}

// Delete takes a key and removes the value stored there if it exists
func (s *Store) Delete(key string) error {
	delete(s.Data, key)
	return nil
}
