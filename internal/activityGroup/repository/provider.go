package querier

import (
	"database/sql"
	"sync"

	"github.com/google/wire"
)

var (
	repo     *activityGroupRepoImpl
	repoOnce sync.Once

	ActivityGroupRepoSet wire.ProviderSet = wire.NewSet(
		ProvideActivityGroupRepo,

		wire.Bind(new(ActivityGroupRepo), new(*activityGroupRepoImpl)),
	)
)

func ProvideActivityGroupRepo(db *sql.DB) (*activityGroupRepoImpl, error) {
	repoOnce.Do(func() {

		repo = &activityGroupRepoImpl{
			db: db,
			Queries: New(db),
		}
	})

	return repo, nil
}
