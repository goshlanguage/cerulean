package lightdb

// Store keeps our data and allows us to pass around a reference to this map.
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
func (s *Store) Put(key string, value string) {
	s.Data[key] = value
}

// Get returns the value for the key if it exists or an empty string
func (s *Store) Get(key string) string {
	if val, ok := s.Data[key]; ok {
		return val
	}
	return ""
}

// Delete takes a key and removes the value stored there if it exists
func (s *Store) Delete(key string) {
	delete(s.Data, key)
}
