package dtos

type CreateTodoRequest struct {
	Title           string `json:"title" validate:"required"`
	ActivityGroupID int    `json:"activity_group_id" validate:"required,gte=1,int"`
	IsActive        bool   `json:"is_active" validate:"required,oneof=true false"`
}

type UpdateTodoRequest struct {
	Title           string `json:"title" validate:"required"`
	ActivityGroupID int    `json:"activity_group_id" validate:"required,gte=1,int"`
	IsActive        bool   `json:"is_active" validate:"required,oneof=true false"`
	Priority        string `json:"priority" validate:"required,oneof=low medium high very-high very-low"`
}
