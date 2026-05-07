// ! main store struct + constructor
package store

import "sync"

type Store struct {
	kv	map[string]string
	queues map[string][]string
	mu sync.RWMutex // RWMutex : many reads and fewer writes
}


func New() *Store{
	return &Store{
		kv: make(map[string]string),
		queues: make(map[string][]string),
		// mu: sync.RWMutex{},  // zero value of RWMutex is ready to use
	}
}