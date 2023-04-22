package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kholiqcode/go-todolist/utils"
)

func SetupMiddleware(route *chi.Mux, config *utils.BaseConfig) {
	if config.Environment == utils.LOCAL || config.Environment == utils.TEST {
		route.Use(Recovery)
	} else {
		route.Use(middleware.RequestID)
		route.Use(middleware.RealIP)
		route.Use(middleware.Logger)
		route.Use(middleware.Timeout(60 * time.Second))

		route.Use(Recovery)
	}

	route.NotFound(func(w http.ResponseWriter, r *http.Request) {
		utils.GenerateJsonResponse(w, nil, http.StatusNotFound, "Not Found")
	})
}
