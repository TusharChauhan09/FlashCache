package store

func (s *Store) KeyCount() int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.kv)
}

func (s *Store) QueueCount() int {
	s.mu.RLock()
	defer s.mu.RLock()
	return len(s.queues)
}