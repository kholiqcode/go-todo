package querier

import (
	"database/sql"
)

type TodoRepo interface {
	Querier

	WithTx(tx *sql.Tx) Querier
	GetDB() *sql.DB
}

type todoRepoImpl struct {
	db *sql.DB
	*Queries
}

func NewTodoRepo(db *sql.DB) TodoRepo {
	return &todoRepoImpl{
		db:      db,
		Queries: New(db),
	}
}

func (r *todoRepoImpl) WithTx(tx *sql.Tx) Querier {
	return &Queries{
		db: tx,
	}
}

func (r *todoRepoImpl) GetDB() *sql.DB {
	return r.db
}
