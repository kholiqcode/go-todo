package v1

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kholiqcode/go-todolist/internal/activityGroup/service"
	"github.com/kholiqcode/go-todolist/utils"
)

type ActivityGroupHandler interface {
	getActivityGroups(w http.ResponseWriter, r *http.Request)
	MapRoutes()
}

type activityGroupHandlerImpl struct {
	route            *chi.Mux
	activityGroupSvc service.ActivityGroupService
}

func (h *activityGroupHandlerImpl) getActivityGroups(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()

	activityGroupsResp, err := h.activityGroupSvc.FindAll(ctx)
	utils.LogAndPanicIfError(err, "failed to get activity groups")

	utils.GenerateJsonResponse(w, activityGroupsResp, 200, "OK")
}
