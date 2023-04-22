package querier

import (
	"database/sql"
	"sync"

	"github.com/google/wire"
)

var (
	repo     *todoRepoImpl
	repoOnce sync.Once

	TodoRepoSet wire.ProviderSet = wire.NewSet(
		ProvideTodoRepo,

		wire.Bind(new(TodoRepo), new(*todoRepoImpl)),
	)
)

func ProvideTodoRepo(db *sql.DB) (*todoRepoImpl, error) {
	repoOnce.Do(func() {

		repo = &todoRepoImpl{
			db: db,
			Queries: New(db),
		}
	})

	return repo, nil
}
