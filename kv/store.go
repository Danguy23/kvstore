package kv

type Store struct { //defines multi-field type with struct //store is an in-memory key-value store
	items map[string]string
}

func New() *Store {
	return &Store{items: make(map[string]string)}
}

func (s *Store) Get(key string) (string, bool) {
	v, ok := s.items[key]
	return v, ok
}

func (s*Store) Put(key, value string) {
	s.items[key] = value
}
func (s*Store) Delete(key string) {
	delete(s.items, key)
}