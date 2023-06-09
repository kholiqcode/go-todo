package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kholiqcode/go-todolist/internal/todo/dtos"
	"github.com/kholiqcode/go-todolist/internal/todo/service"
	"github.com/kholiqcode/go-todolist/utils"
)

type TodoHandler interface {
	getTodos(w http.ResponseWriter, r *http.Request)
	getTodo(w http.ResponseWriter, r *http.Request)
	createTodo(w http.ResponseWriter, r *http.Request)
	updateTodo(w http.ResponseWriter, r *http.Request)
	deleteTodo(w http.ResponseWriter, r *http.Request)
	MapRoutes()
}
type todoHandlerImpl struct {
	route   *chi.Mux
	todoSvc service.TodoService
}

func (h *todoHandlerImpl) getTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	activityGroupID := utils.ValidateQueryParamInt(r, "activity_group_id")

	request := dtos.GetTodosRequest{
		ActivityGroupID: int32(activityGroupID),
	}

	activityGroupsResp, err := h.todoSvc.FindAll(ctx, request)
	utils.PanicIfError(err)

	utils.GenerateJsonResponse(w, activityGroupsResp, 200, "Success")
}

func (h *todoHandlerImpl) getTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := utils.ValidateUrlParamInt(r, "id")

	todoResp, err := h.todoSvc.FindByID(ctx, int32(id))
	utils.PanicIfError(err)

	utils.GenerateJsonResponse(w, todoResp, 200, "Success")
}

func (h *todoHandlerImpl) createTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request dtos.CreateTodoRequest

	utils.ValidateBodyPayload(r.Body, &request)

	if request.ActivityGroupID == 0 {
		utils.PanicAppError("activity_group_id cannot be null", 400)
	}

	todoResp, err := h.todoSvc.Store(ctx, request)
	utils.PanicIfError(err)

	utils.GenerateJsonResponse(w, todoResp, 201, "Success")
}

func (h *todoHandlerImpl) updateTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var request dtos.UpdateTodoRequest

	id := utils.ValidateUrlParamInt(r, "id")

	utils.ValidateBodyPayload(r.Body, &request)

	todoResp, err := h.todoSvc.Update(ctx, int32(id), request)
	utils.PanicIfError(err)

	utils.GenerateJsonResponse(w, todoResp, 200, "Success")
}

func (h *todoHandlerImpl) deleteTodo(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id := utils.ValidateUrlParamInt(r, "id")

	err := h.todoSvc.Delete(ctx, int32(id))
	utils.PanicIfError(err)

	any := struct{}{}

	utils.GenerateJsonResponse(w, any, 200, "Success")
}
