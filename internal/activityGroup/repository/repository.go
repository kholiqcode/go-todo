package querier

import (
	"database/sql"
)

type ActivityGroupRepo interface {
	Querier

	WithTx(tx *sql.Tx) Querier
	GetDB() *sql.DB
}

type activityGroupRepoImpl struct {
	db *sql.DB
	*Queries
}

func NewActivityGroupRepo(db *sql.DB) ActivityGroupRepo {
	return &activityGroupRepoImpl{
		db:      db,
		Queries: New(db),
	}
}

func (r *activityGroupRepoImpl) WithTx(tx *sql.Tx) Querier {
	return &Queries{
		db: tx,
	}
}

func (r *activityGroupRepoImpl) GetDB() *sql.DB {
	return r.db
}
