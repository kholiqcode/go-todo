package v1

import (
	"net/http"
	"strconv"

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

	ctx := r.Context()

	activityGroupsResp, err := h.activityGroupSvc.FindAll(ctx)
	utils.LogAndPanicIfError(err, "failed to get activity groups")

	utils.GenerateJsonResponse(w, activityGroupsResp, 200, "Success")
}

func (h *activityGroupHandlerImpl) getActivityGroup(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	id := chi.URLParam(r, "id")

	idInt, err := strconv.Atoi(id)

	utils.LogAndPanicIfError(err, "failed to convert id to int")

	activityGroupResp, err := h.activityGroupSvc.FindByID(ctx, int32(idInt))
	if err != nil {
		utils.GenerateJsonResponse(w, nil, 404, "Not Found")
		return
	}

	utils.GenerateJsonResponse(w, activityGroupResp, 200, "Success")
}
