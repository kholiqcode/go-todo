package app

import (
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
	v1_activityGroup "github.com/kholiqcode/go-todolist/internal/activityGroup/delivery/http/v1"
	v1_todo "github.com/kholiqcode/go-todolist/internal/todo/delivery/http/v1"
	"github.com/kholiqcode/go-todolist/utils"
)

var (
	httpServer *httpServerImpl
	httpOnce   sync.Once

	AppSet wire.ProviderSet = wire.NewSet(
		ProvideHttpServer,

		wire.Bind(new(HttpServer), new(*httpServerImpl)),
	)
)

func ProvideHttpServer(route *chi.Mux, config *utils.BaseConfig, activityHandler v1_activityGroup.ActivityGroupHandler, todoHandler v1_todo.TodoHandler) (*httpServerImpl, error) {
	httpOnce.Do(func() {

		httpServer = &httpServerImpl{
			route:           route,
			config:          config,
			activityHandler: activityHandler,
			todoHandler:     todoHandler,
			startAt:         time.Now(),
		}
	})

	return httpServer, nil
}
