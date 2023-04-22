package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kholiqcode/go-todolist/internal/activityGroup/dtos"
	"github.com/kholiqcode/go-todolist/internal/activityGroup/service"
	"github.com/kholiqcode/go-todolist/utils"
)

type ActivityGroupHandler interface {
	getActivityGroups(w http.ResponseWriter, r *http.Request)
	getActivityGroup(w http.ResponseWriter, r *http.Request)
	createActivityGroup(w http.ResponseWriter, r *http.Request)
	updateActivityGroup(w http.ResponseWriter, r *http.Request)
	deleteActivityGroup(w http.ResponseWriter, r *http.Request)
	MapRoutes()
}

type activityGroupHandlerImpl struct {
	route            *chi.Mux
	activityGroupSvc service.ActivityGroupService
}

func (h *activityGroupHandlerImpl) getActivityGroups(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	activityGroupsResp, err := h.activityGroupSvc.FindAll(ctx)
	utils.PanicIfError(err)

	utils.GenerateJsonResponse(w, activityGroupsResp, 200, "Success")
}

func (h *activityGroupHandlerImpl) getActivityGroup(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	id := utils.ValidateUrlParamInt(r, "id")

	activityGroupResp, err := h.activityGroupSvc.FindByID(ctx, int32(id))
	utils.PanicIfError(err)

	utils.GenerateJsonResponse(w, activityGroupResp, 200, "Success")
}

func (h *activityGroupHandlerImpl) createActivityGroup(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var request dtos.CreateActivityGroupRequest

	utils.ValidateBodyPayload(r.Body, &request)

	activityGroupResp, err := h.activityGroupSvc.Store(ctx, request)
	utils.PanicIfError(err)

	utils.GenerateJsonResponse(w, activityGroupResp, 201, "Success")
}

func (h *activityGroupHandlerImpl) updateActivityGroup(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var request dtos.UpdateActivityGroupRequest

	id := utils.ValidateUrlParamInt(r, "id")

	utils.ValidateBodyPayload(r.Body, &request)

	activityGroupResp, err := h.activityGroupSvc.Update(ctx, int32(id), request)
	utils.PanicIfError(err)

	utils.GenerateJsonResponse(w, activityGroupResp, 200, "Success")
}

func (h *activityGroupHandlerImpl) deleteActivityGroup(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	id := utils.ValidateUrlParamInt(r, "id")

	err := h.activityGroupSvc.Delete(ctx, int32(id))
	utils.PanicIfError(err)

	utils.GenerateJsonResponse(w, nil, 204, "Success")
}
