package db

import (
	"context"
	"sync"

	"gorm.io/gorm"
)

// temp name
type SQLManager struct {
	mu          *sync.Mutex
	conns       map[string]*gorm.DB
	trxPrefixes map[string]map[string]bool
}

func NewSQLManager(conns map[string]*gorm.DB) *SQLManager {
	return &SQLManager{conns: conns, mu: &sync.Mutex{}, trxPrefixes: map[string]map[string]bool{}}
}

func (m *SQLManager) Begin(ctx context.Context) error {
	return nil
}
func (m *SQLManager) Commit(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if ctx.Value("trx_id") == nil {
		return nil
	}

	trxID, ok := ctx.Value("trx_id").(string)
	if !ok {
		return nil
	}

	if _, exist := m.trxPrefixes[trxID]; !exist {
		return nil
	}

	for prefix := range m.trxPrefixes[trxID] {
		conn, exist := m.conns[prefix+"_"+trxID]
		if !exist {
			continue
		}

		err := conn.Commit().Error
		if err != nil {
			m.Rollback(ctx)
			return err
		}
		delete(m.conns, prefix+"_"+trxID)
	}

	return err
}

func (m *SQLManager) Rollback(ctx context.Context) error {

	m.mu.Lock()
	defer m.mu.Unlock()

	if ctx.Value("trx_id") == nil {
		return nil
	}

	trxID, ok := ctx.Value("trx_id").(string)
	if !ok {
		return nil
	}

	_, exist := m.trxPrefixes[trxID]
	if !exist {
		return nil
	}
	for prefix := range m.trxPrefixes[trxID] {
		conn, exist := m.conns[prefix+"_"+trxID]
		if !exist {
			continue
		}
		conn.Rollback()
		delete(m.conns, prefix+"_"+trxID)
	}

	return err
}

// GetConn gets a connection by name or starts a transaction if given a trx_id ( from context ) .
func (m *SQLManager) GetConn(ctx context.Context, connName string) *gorm.DB {
	m.mu.Lock()
	defer m.mu.Unlock()

	trxID, ok := ctx.Value("trx_id").(string)
	if !ok || trxID == "" {
		return m.conns[connName]
	}

	_, exist := m.trxPrefixes[trxID]
	if !exist {
		m.trxPrefixes[trxID] = map[string]bool{}
		m.conns[connName+"_"+trxID] = m.conns[connName].Begin()
	}

	m.trxPrefixes[trxID][connName] = true

	return m.conns[connName+"_"+trxID]
}

func (m *SQLManager) HasTransaction(ctx context.Context) bool {
	trxID, _ := ctx.Value("trx_id").(string)
	return trxID != ""
}
