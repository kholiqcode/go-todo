package dtos

import querier "github.com/kholiqcode/go-todolist/internal/todo/repository"

func ToTodoResponse(todo querier.Todo) TodoResponse {
	return TodoResponse{
		ID:              todo.ID,
		Title:           todo.Title,
		ActivityGroupID: todo.ActivityGroupID,
		IsActive:        todo.IsActive,
		Priority:        todo.Priority,
		CreatedAt:       &todo.CreatedAt.Time,
		UpdatedAt:       &todo.UpdatedAt.Time,
	}
}
