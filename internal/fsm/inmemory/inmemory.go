package inmemory

import "sync"

type InMemoryStateManager struct {
	mu     sync.RWMutex
	states map[int64]string
}

func NewInMemoryStateManager() *InMemoryStateManager {
	return &InMemoryStateManager{
		states: make(map[int64]string),
	}
}

func (m *InMemoryStateManager) Set(userID int64, state string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.states[userID] = state
}

func (m *InMemoryStateManager) Get(userID int64) string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.states[userID]
}

func (m *InMemoryStateManager) Clear(userID int64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.states, userID)
}
