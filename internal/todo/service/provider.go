package service

import (
	"sync"

	"github.com/google/wire"
	querier "github.com/kholiqcode/go-todolist/internal/todo/repository"
)

var (
	svc     *todoServiceImpl
	svcOnce sync.Once

	TodoServiceSet wire.ProviderSet = wire.NewSet(
		ProvideTodoService,

		wire.Bind(new(TodoService), new(*todoServiceImpl)),
	)
)

func ProvideTodoService(querier querier.TodoRepo) (*todoServiceImpl, error) {
	svcOnce.Do(func() {
		svc = &todoServiceImpl{
			repo: querier,
		}
	})

	return svc, nil
}
