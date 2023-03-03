package db

import (
	"context"
	"errors"

	"github.com/crewdible/go-lib/stringlib"
)

type (
	contextTranasction struct {
		trxs []Transaction
	}

	Transaction interface {
		Begin(ctx context.Context) error
		Commit(ctx context.Context) error
		Rollback(ctx context.Context) error
	}

	ContextTransaction interface {
		Begin(ctx context.Context) (context.Context, error)
		Commit(ctx context.Context) (context.Context, error)
		Rollback(ctx context.Context) (context.Context, error)
		RollbackIfNotCommited(ctx context.Context) (context.Context, error)
	}
)

func NewContextTransaction(trxs ...Transaction) ContextTransaction {
	return &contextTranasction{trxs: trxs}
}

func (trx *contextTranasction) Begin(ctx context.Context) (context.Context, error) {
	ctx = context.WithValue(ctx, "trx_id", stringlib.GenerateRandString(6))
	for _, da := range trx.trxs {
		err := da.Begin(ctx)
		if err != nil {
			return ctx, err
		}
	}
	return ctx, nil
}

func (trx *contextTranasction) Commit(ctx context.Context) (context.Context, error) {
	if ctx.Value("trx_id") == nil {
		return ctx, errors.New("no transaction to commit")
	}
	for _, da := range trx.trxs {
		err := da.Commit(ctx)
		if err != nil {
			return ctx, err
		}
	}

	ctx = context.WithValue(ctx, "trx_id", nil)
	return ctx, nil
}

func (trx *contextTranasction) Rollback(ctx context.Context) (context.Context, error) {
	if ctx.Value("trx_id") == nil {
		return ctx, errors.New("no transaction to rollback")
	}
	for _, da := range trx.trxs {
		err := da.Rollback(ctx)
		if err != nil {
			return ctx, err
		}
	}
	ctx = context.WithValue(ctx, "trx_id", nil)
	return ctx, nil
}

func (trx *contextTranasction) RollbackIfNotCommited(ctx context.Context) (context.Context, error) {
	for _, da := range trx.trxs {
		err := da.Rollback(ctx)
		if err != nil {
			return ctx, err
		}
	}
	ctx = context.WithValue(ctx, "trx_id", nil)
	return ctx, nil
}
