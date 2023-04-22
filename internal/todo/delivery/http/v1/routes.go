package v1

import "github.com/go-chi/chi/v5"

func (h *todoHandlerImpl) MapRoutes() {

	h.route.Route("/todo-items", func(r chi.Router) {
		r.Get("/", h.getTodos)
		r.Get("/{id}", h.getTodo)
		r.Post("/", h.createTodo)
		r.Patch("/{id}", h.updateTodo)
	})
}
