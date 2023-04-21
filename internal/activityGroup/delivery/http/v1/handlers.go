package v1

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/kholiqcode/go-todolist/utils"
)

type ActivityGroupHandler interface {
	health(w http.ResponseWriter, r *http.Request)
	MapRoutes()
}

type activityGroupHandlerImpl struct {
	route *chi.Mux
}

func (h *activityGroupHandlerImpl) health(w http.ResponseWriter, r *http.Request) {
	utils.GenerateJsonResponse(w, nil, http.StatusOK, "OK")
}
