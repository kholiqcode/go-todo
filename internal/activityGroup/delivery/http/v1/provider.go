package v1

import (
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
)

var (
	hdl     *activityGroupHandlerImpl
	hdlOnce sync.Once

	ActivityGroupSet wire.ProviderSet = wire.NewSet(
		ProvideHttpServer,

		wire.Bind(new(ActivityGroupHandler), new(*activityGroupHandlerImpl)),
	)
)

func ProvideHttpServer(route *chi.Mux) (*activityGroupHandlerImpl, error) {
	hdlOnce.Do(func() {

		hdl = &activityGroupHandlerImpl{
			route: route,
		}
	})

	return hdl, nil
}
