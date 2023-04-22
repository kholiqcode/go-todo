package dtos

import (
	"time"

	querier "github.com/kholiqcode/go-todolist/internal/todo/repository"
)

type TodoResponse struct {
	ID              int32                 `json:"id"`
	Title           string                `json:"title"`
	ActivityGroupID int32                 `json:"activity_group_id"`
	IsActive        bool                  `json:"is_active"`
	Priority        querier.TodosPriority `json:"priority"`
	CreatedAt       *time.Time            `json:"createdAt"`
	UpdatedAt       *time.Time            `json:"updatedAt"`
}
