package dtos

type CreateTodoRequest struct {
	Title           string `json:"title" validate:"required"`
	ActivityGroupID int    `json:"activity_group_id" validate:"required,gte=1"`
	IsActive        bool   `json:"is_active" validate:"required"`
}

type UpdateTodoRequest struct {
	Title string `json:"title" validate:"required"`
}

type GetTodosRequest struct {
	ActivityGroupID int32 `json:"activity_group_id" validate:"required,gte=1,int"`
}
