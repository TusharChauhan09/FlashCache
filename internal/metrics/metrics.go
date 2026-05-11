// ! Metrics struct
package metrics

import (
	"sync"
	"time"
)

type Metrics struct{
	Connections int
	Requests int
	StartTime time.Time

	mu sync.RWMutex
}

func New() *Metrics{
	return &Metrics{
		StartTime: time.Now(),
	}
}

