package storage

type Store struct{
	data map[string]string
}

func NewStore() *Store{
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Put(key, value string){
	s.data[key] = value
}

