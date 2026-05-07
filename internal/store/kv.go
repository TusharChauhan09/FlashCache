// ! SET / GET / DEL / EXISTS
package store

func (s *Store) Set(key, value string ){
	s.mu.Lock()
	defer s.mu.Unlock()

	s.kv[key] = value
}

func (s *Store) Get(key string) (string,error){
	s.mu.RLock()
	defer s.mu.RUnlock()

	value, exist := s.kv[key]  // ok==exist
	if !exist {
		return "",ErrKeyNotFound
	} 

	return value, nil
}

func (s *Store) Delete(key string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_ , exist := s.kv[key]
	if !exist{
		return false
	}
	
	delete(s.kv,key)
	return true
}

func (s *Store) Exist(key string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	_ , exist := s.kv[key]

	return exist
}