package v1

func (h *activityGroupHandlerImpl) MapRoutes() {

	h.route.Mount("/activity-groups", h.route)

	h.route.Get("/", h.getActivityGroups)
	h.route.Get("/{id}", h.getActivityGroup)
	h.route.Post("/", h.createActivityGroup)

}
