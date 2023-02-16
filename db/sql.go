package db

import (
	"context"
	"sync"

	"gorm.io/gorm"
)

// temp name
type SQLManager struct {
	mu    *sync.Mutex
	conns map[string]*gorm.DB
}

func NewSQLManager(conns map[string]*gorm.DB) *SQLManager {
	return &SQLManager{conns: conns, mu: &sync.Mutex{}}
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
	err := m.conns[trxID].Commit().Error
	delete(m.conns, trxID)

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

	err := m.conns[trxID].Rollback().Error
	delete(m.conns, trxID)
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

	if _, exist := m.conns[trxID]; !exist {
		m.conns[trxID] = m.conns[connName].Begin()
	}

	return m.conns[trxID]
}
