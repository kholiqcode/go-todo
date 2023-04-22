package v1

import "github.com/go-chi/chi/v5"

func (h *activityGroupHandlerImpl) MapRoutes() {

	h.route.Route("/activity-groups", func(r chi.Router) {
		r.Get("/", h.getActivityGroups)
		r.Get("/{id}", h.getActivityGroup)
		r.Post("/", h.createActivityGroup)
		r.Patch("/{id}", h.updateActivityGroup)
		r.Delete("/{id}", h.deleteActivityGroup)
	})

}
