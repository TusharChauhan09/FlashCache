// !  PUSH / POP / LEN
package store

func (s *Store) Push(queueName , value string){
	s.mu.Lock()
	defer s.mu.Unlock()

	s.queues[queueName] = append(s.queues[queueName], value)
}

func (s *Store) Pop(queueName string) (string, error){
	s.mu.Lock()
	defer s.mu.Unlock()

	queue, exists := s.queues[queueName]
	if !exists || len(queue) == 0 {
		return "", ErrQueueEmpty
	}

	value := queue[0]
	s.queues[queueName] = queue[1:]

	return value, nil
}

func (s *Store) Len (queueName string) int {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return len(s.queues[queueName])
}