// ! STATS command logic
package metrics

import (
	"time"
	"fmt"
)

func (m *Metrics) IncrementConnections(){
	m.mu.Lock()
	defer m.mu.Unlock()

	m.Connections++;
}

func (m *Metrics) IncrementRequests() {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.Requests++
}

func (m *Metrics) GetConnections() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.Connections
}

func (m *Metrics) GetRequests() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.Requests
}

func (m *Metrics) GetUptime() time.Duration {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return time.Since(m.StartTime)
}

func (m *Metrics) Stats(keys,queues int) string{
	return fmt.Sprintf(
		"connections=%d requests=%d keys=%d queues=%d uptime=%s",
		m.GetConnections(),
		m.GetRequests(),
		keys,
		queues,
		m.GetUptime().Round(time.Second),
	)
}


