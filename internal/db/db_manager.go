package db

import (
	"fmt"
	"path/filepath"
	"sync"
)

// DBManager manages multiple RocksDB instances
type DBManager struct {
	baseDir string
	dbs     map[string]*RocksDB
	mu      sync.RWMutex
}

// NewDBManager creates a new database manager
func NewDBManager(baseDir string) *DBManager {
	return &DBManager{
		baseDir: baseDir,
		dbs:     make(map[string]*RocksDB),
	}
}

// GetDB returns an existing database or creates a new one
func (m *DBManager) GetDB(name string) (*RocksDB, error) {
	if name == "" {
		return nil, fmt.Errorf("database name cannot be empty")
	}

	m.mu.RLock()
	if db, exists := m.dbs[name]; exists {
		m.mu.RUnlock()
		return db, nil
	}
	m.mu.RUnlock()

	// If we get here, we need to create a new database
	m.mu.Lock()
	defer m.mu.Unlock()

	// Double-check if another goroutine created the database
	if db, exists := m.dbs[name]; exists {
		return db, nil
	}

	// Create new database directory
	dbPath := filepath.Join(m.baseDir, name)
	db, err := NewRocksDB(dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create database %s: %w", name, err)
	}

	m.dbs[name] = db
	return db, nil
}

// Close closes all database instances
func (m *DBManager) Close() {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, db := range m.dbs {
		db.Close()
	}
	m.dbs = make(map[string]*RocksDB)
}
