package session

import (
	"errors"
	"sync"
	"time"
)

var ErrSessionNotFound = errors.New("session not found")

type CustomerSession struct {
	SessionID    string
	StartedAt    time.Time
	LastActiveAt time.Time
	Context      map[string]interface{}
	Messages     []map[string]interface{}
}

type Manager struct {
	sessions map[string]*CustomerSession
	mu       sync.RWMutex
}

func NewManager() *Manager {
	return &Manager{
		sessions: make(map[string]*CustomerSession),
	}
}

func (m *Manager) CreateSession(id string) *CustomerSession {
	m.mu.Lock()
	defer m.mu.Unlock()

	session := &CustomerSession{
		SessionID:    id,
		StartedAt:    time.Now(),
		LastActiveAt: time.Now(),
		Context:      make(map[string]interface{}),
		Messages:     make([]map[string]interface{}, 0),
	}
	m.sessions[id] = session
	return session
}

func (m *Manager) GetSession(id string) (*CustomerSession, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if session, ok := m.sessions[id]; ok {
		return session, nil
	}
	return nil, ErrSessionNotFound
}

func (m *Manager) UpdateSession(id string, fn func(*CustomerSession)) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if session, ok := m.sessions[id]; ok {
		fn(session)
		session.LastActiveAt = time.Now()
		return nil
	}
	return ErrSessionNotFound
}

