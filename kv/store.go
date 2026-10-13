package kv

type Store struct { //defines multi-field type with struct //store is an in-memory key-value store
	items map[string]string
}

func New() *Store { //creates a new empty map instance
	return &Store{items: make(map[string]string)}
}

func (s *Store) Get(key string) (string, bool) { //retrieves value using key
	v, ok := s.items[key]
	return v, ok
}

func (s*Store) Put(key, value string) { //adds or updates a key-value pair in the store
	s.items[key] = value
}
func (s*Store) Delete(key string) { //removes a key-value pair from the store
	delete(s.items, key)
}