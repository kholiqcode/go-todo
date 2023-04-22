package v1

import (
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
	"github.com/kholiqcode/go-todolist/internal/todo/service"
)

var (
	hdl     *todoHandlerImpl
	hdlOnce sync.Once

	TodoHandlerSet wire.ProviderSet = wire.NewSet(
		ProvideTodoHandler,

		wire.Bind(new(TodoHandler), new(*todoHandlerImpl)),
	)
)

func ProvideTodoHandler(route *chi.Mux, todoSvc service.TodoService) (*todoHandlerImpl, error) {
	hdlOnce.Do(func() {

		hdl = &todoHandlerImpl{
			route:   route,
			todoSvc: todoSvc,
		}
	})

	return hdl, nil
}
