package v1

func (h *activityGroupHandlerImpl) MapRoutes() {
	h.route.Mount("/activity-groups", h.route)

	h.route.Get("/test", h.getActivityGroups)
}
