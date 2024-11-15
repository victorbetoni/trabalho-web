package uow

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"victorbetoni/trabalho-web/internal/infra/util"
)

type UowInterface interface {
	Register(string, RepositoryFactory)
	GetRepository(context.Context, string) (interface{}, error)
	Do(context.Context, func(uow *Uow) util.Response) util.Response
	CommitOrRollback(util.Response) util.Response
	Rollback(util.Response) util.Response
	UnRegister(string)
}

var current *Uow

type RepositoryFactory func(tx *sql.Tx) interface{}

type Uow struct {
	Db           *sql.DB
	Tx           *sql.Tx
	mu           sync.Mutex
	Repositories map[string]RepositoryFactory
}

func Current() *Uow {
	return current
}

func NewUow(ctx context.Context, db *sql.DB) (*Uow, error) {
	current = &Uow{
		Repositories: make(map[string]RepositoryFactory),
		Db:           db,
	}
	return current, nil
}

func (u *Uow) Register(name string, fc RepositoryFactory) {
	u.Repositories[name] = fc
}

func (u *Uow) GetRepository(ctx context.Context, name string) (interface{}, error) {
	if u.Tx == nil {
		tx, err := u.Db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		u.Tx = tx
	}
	repo := u.Repositories[name](u.Tx)
	return repo, nil
}

func (u *Uow) Do(ctx context.Context, fn func(uow *Uow) util.Response) util.Response {
	u.mu.Lock()
	if u.Tx != nil {
		return util.InternalServerErr(errors.New("transaction already started"))
	}
	tx, err := u.Db.BeginTx(ctx, nil)
	if err != nil {
		return util.InternalServerErr(err)
	}
	u.Tx = tx
	res := fn(u)
	if !res.Ok() {
		return u.Rollback(res)
	}
	return u.CommitOrRollback(res)
}

func (u *Uow) CommitOrRollback(res util.Response) util.Response {
	err := u.Tx.Commit()
	if err != nil {
		if resp := u.Rollback(util.InternalServerErr(err)); !resp.Ok() {
			u.mu.Unlock()
			return util.InternalServerErr(fmt.Errorf("commit error: %s, rollback error: %s", err, resp.Message))
		}
		u.mu.Unlock()
		return util.InternalServerErr(err)
	}
	u.Tx = nil
	u.mu.Unlock()
	return res
}

func (u *Uow) Rollback(res util.Response) util.Response {
	if u.Tx == nil {
		u.mu.Unlock()
		return util.InternalServerErr(errors.New("no transaction to rollback"))
	}
	err := u.Tx.Rollback()
	if err != nil {
		u.mu.Unlock()
		return util.InternalServerErr(err)
	}
	u.Tx = nil
	u.mu.Unlock()
	return res
}

func (u *Uow) UnRegister(name string) {
	delete(u.Repositories, name)
}
