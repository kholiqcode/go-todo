package v1

import (
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/google/wire"
	"github.com/kholiqcode/go-todolist/internal/activityGroup/service"
)

var (
	hdl     *activityGroupHandlerImpl
	hdlOnce sync.Once

	ActivityGroupHandlerSet wire.ProviderSet = wire.NewSet(
		ProvideActivityGroupHandler,

		wire.Bind(new(ActivityGroupHandler), new(*activityGroupHandlerImpl)),
	)
)

func ProvideActivityGroupHandler(route *chi.Mux, activityGroupSvc service.ActivityGroupService) (*activityGroupHandlerImpl, error) {
	hdlOnce.Do(func() {

		hdl = &activityGroupHandlerImpl{
			route:            route,
			activityGroupSvc: activityGroupSvc,
		}
	})

	return hdl, nil
}
