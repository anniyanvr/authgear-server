package db

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/authgear/authgear-server/pkg/util/errorutil"
	"github.com/authgear/authgear-server/pkg/util/log"
	"github.com/authgear/authgear-server/pkg/util/uuid"
)

type hookHandleContextKeyType struct{}

var hookHandleContextKey = hookHandleContextKeyType{}

type hookHandleContextValue struct {
	ConnLike ConnLike
	Hooks    []TransactionHook
}

type HookHandle struct {
	Pool              *Pool
	ConnectionOptions ConnectionOptions
	Logger            *log.Logger
}

func hookHandleContextWithValue(ctx context.Context, value *hookHandleContextValue) context.Context {
	return context.WithValue(ctx, hookHandleContextKey, value)
}

func hookHandleContextGetValue(ctx context.Context) (*hookHandleContextValue, bool) {
	v, ok := ctx.Value(hookHandleContextKey).(*hookHandleContextValue)
	if !ok {
		return nil, false
	}
	return v, true
}

func mustHookHandleContextGetValue(ctx context.Context) *hookHandleContextValue {
	v, ok := hookHandleContextGetValue(ctx)
	if !ok {
		panic(fmt.Errorf("hook-handle: transaction not started"))
	}
	return v
}

func mustGetConnLike(ctx context.Context) ConnLike {
	return mustHookHandleContextGetValue(ctx).ConnLike
}

var _ Handle = (*HookHandle)(nil)

func NewHookHandle(pool *Pool, opts ConnectionOptions, lf *log.Factory) *HookHandle {
	return &HookHandle{
		Pool:              pool,
		ConnectionOptions: opts,
		Logger:            lf.New("db-handle"),
	}
}

func (h *HookHandle) UseHook(ctx context.Context, hook TransactionHook) {
	v, ok := hookHandleContextGetValue(ctx)
	if !ok {
		panic(fmt.Errorf("hook-handle: transaction not started"))
	}

	v.Hooks = append(v.Hooks, hook)
}

// WithTx commits if do finishes without error and rolls back otherwise.
// WithTx is reentrant, meaning that you can call WithTx even when a previous WithTx does not finish yet.
// Normally you should not call WithTx within a WithTx, but there is a legit use case.
//
//	// Assume ctx is a http.Request context.
//	h.WithTx(ctx, func(ctx context.Context) error {
//		// ctx here is associated with a *sql.Tx (Tx1)
//		go func() {
//			// ctx is detached from the http.Request context.
//			ctx = ctx.WithCancel(ctx)
//			h.WithTx(ctx, func(ctx context.Context) error {
//				// ctx is associated with a *sqlTx (Tx2)
//			})
//		}()
//	})
func (h *HookHandle) WithTx(ctx context.Context, do func(ctx context.Context) error) (err error) {
	id := uuid.New()
	logger := h.Logger.WithField("debug_id", id)
	db, err := h.openDB()
	if err != nil {
		return
	}

	conn, err := db.Conn(ctx)
	if err != nil {
		err = fmt.Errorf("hook-handle: failed to acquire connection: %w", err)
		return
	}
	logger.Debug("acquire connection")

	tx, err := h.beginTx(ctx, logger, conn)
	if err != nil {
		return
	}

	ctx = hookHandleContextWithValue(ctx, &hookHandleContextValue{
		ConnLike: tx,
	})

	shouldRunDidCommitHooks := false

	defer func() {
		if shouldRunDidCommitHooks {
			for _, hook := range mustHookHandleContextGetValue(ctx).Hooks {
				hook.DidCommitTx(ctx)
			}
		}
	}()

	defer func() {
		closeErr := conn.Close()
		if closeErr != nil && !errors.Is(closeErr, sql.ErrConnDone) {
			logger.WithError(closeErr).Error("failed to close connection")
		} else {
			logger.Debug("close connection")
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			_ = rollbackTx(logger, tx)
			panic(r)
		} else if err != nil {
			_ = rollbackTx(logger, tx)
		} else {
			err = commitTx(ctx, logger, tx, mustHookHandleContextGetValue(ctx).Hooks)
			if err == nil {
				shouldRunDidCommitHooks = true
			}
		}
	}()

	err = do(ctx)
	return
}

// ReadOnly is like WithTx, except that it always rolls back.
func (h *HookHandle) ReadOnly(ctx context.Context, do func(ctx context.Context) error) (err error) {
	id := uuid.New()
	logger := h.Logger.WithField("debug_id", id)
	db, err := h.openDB()
	if err != nil {
		return
	}

	conn, err := db.Conn(ctx)
	if err != nil {
		err = fmt.Errorf("hook-handle: failed to acquire connection: %w", err)
		return
	}
	logger.Debug("acquire connection")

	tx, err := h.beginTx(ctx, logger, conn)
	if err != nil {
		return
	}

	ctx = hookHandleContextWithValue(ctx, &hookHandleContextValue{
		ConnLike: tx,
	})

	shouldRunDidCommitHooks := false

	defer func() {
		if shouldRunDidCommitHooks {
			for _, hook := range mustHookHandleContextGetValue(ctx).Hooks {
				hook.DidCommitTx(ctx)
			}
		}
	}()

	defer func() {
		closeErr := conn.Close()
		if closeErr != nil && !errors.Is(closeErr, sql.ErrConnDone) {
			logger.WithError(closeErr).Error("failed to close connection")
		} else {
			logger.Debug("close connection")
		}
	}()

	defer func() {
		if r := recover(); r != nil {
			_ = rollbackTx(logger, tx)
			panic(r)
		} else if err != nil {
			_ = rollbackTx(logger, tx)
		} else {
			err = rollbackTx(logger, tx)
			if err == nil {
				shouldRunDidCommitHooks = true
			}
		}
	}()

	err = do(ctx)
	return
}

func (h *HookHandle) beginTx(ctx context.Context, logger *log.Logger, conn *sql.Conn) (*sql.Tx, error) {
	// Pass a nil TxOptions to use default isolation level.
	var txOptions *sql.TxOptions
	tx, err := conn.BeginTx(ctx, txOptions)
	if err != nil {
		return nil, fmt.Errorf("hook-handle: failed to begin transaction: %w", err)
	}

	logger.Debug("begin")
	return tx, nil
}

func commitTx(ctx context.Context, logger *log.Logger, tx *sql.Tx, hooks []TransactionHook) error {
	for _, hook := range hooks {
		err := hook.WillCommitTx(ctx)
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = errorutil.WithSecondaryError(err, rbErr)
			}
			return err
		}
	}

	err := tx.Commit()
	if err != nil {
		return fmt.Errorf("hook-handle: failed to commit transaction: %w", err)
	}
	logger.Debug("commit")
	return nil
}

func rollbackTx(logger *log.Logger, tx *sql.Tx) error {
	err := tx.Rollback()
	if err != nil {
		return fmt.Errorf("hook-handle: failed to rollback transaction: %w", err)
	}
	logger.Debug("rollback")

	return nil
}

func (h *HookHandle) openDB() (*sql.DB, error) {
	h.Logger.WithFields(map[string]interface{}{
		"max_open_conns":             h.ConnectionOptions.MaxOpenConnection,
		"max_idle_conns":             h.ConnectionOptions.MaxIdleConnection,
		"conn_max_lifetime_seconds":  h.ConnectionOptions.MaxConnectionLifetime.Seconds(),
		"conn_max_idle_time_seconds": h.ConnectionOptions.IdleConnectionTimeout.Seconds(),
	}).Debug("open database")

	db, err := h.Pool.Open(h.ConnectionOptions)
	if err != nil {
		return nil, fmt.Errorf("hook-handle: failed to connect to database: %w", err)
	}

	return db, nil
}
