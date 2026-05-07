// ! custom errors
package store

import "errors"

var (
	ErrKeyNotFound = errors.New("key not found")
	ErrQueueEmpty = errors.New("queue is empty")
)