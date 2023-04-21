package service

import (
	"sync"

	"github.com/google/wire"
	querier "github.com/kholiqcode/go-todolist/internal/activityGroup/repository"
)

var (
	svc     *activityGroupServiceImpl
	svcOnce sync.Once

	ActivityGroupServiceSet wire.ProviderSet = wire.NewSet(
		ProvideActivityGroupService,

		wire.Bind(new(ActivityGroupService), new(*activityGroupServiceImpl)),
	)
)

func ProvideActivityGroupService(querier querier.ActivityGroupRepo) (*activityGroupServiceImpl, error) {
	svcOnce.Do(func() {
		svc = &activityGroupServiceImpl{
			repo: querier,
		}
	})

	return svc, nil
}
