package v1

import "github.com/go-chi/chi/v5"

func (h *todoHandlerImpl) MapRoutes() {

	h.route.Route("/todo-items", func(r chi.Router) {
		r.Get("/", h.getTodos)
	})
}
