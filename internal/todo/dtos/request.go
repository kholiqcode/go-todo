package dtos

type CreateTodoRequest struct {
	Title           string `json:"title" validate:"required"`
	ActivityGroupID int    `json:"activity_group_id"`
	IsActive        bool   `json:"is_active"`
}

type UpdateTodoRequest struct {
	Title    string `json:"title"`
	IsActive bool   `json:"is_active"`
}

type GetTodosRequest struct {
	ActivityGroupID int32 `json:"activity_group_id" validate:"required,gte=1,int"`
}
