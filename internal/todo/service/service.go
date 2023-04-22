package service

import (
	"context"

	"github.com/kholiqcode/go-todolist/internal/todo/dtos"
	querier "github.com/kholiqcode/go-todolist/internal/todo/repository"
	"github.com/kholiqcode/go-todolist/utils"
)

type TodoService interface {
	FindAll(ctx context.Context, request dtos.GetTodosRequest) ([]dtos.TodoResponse, error)
	FindByID(ctx context.Context, id int32) (*dtos.TodoResponse, error)
	// Store(ctx context.Context, request dtos.CreateTodoRequest) (*dtos.TodoResponse, error)
	// Update(ctx context.Context, id int32, request dtos.UpdateTodoRequest) (*dtos.TodoResponse, error)
	// Delete(ctx context.Context, id int32) error
}

type todoServiceImpl struct {
	repo querier.TodoRepo
}

func (s *todoServiceImpl) FindAll(ctx context.Context, request dtos.GetTodosRequest) ([]dtos.TodoResponse, error) {

	params := querier.ListTodosParams{
		SearchField: "activity_group_id",
		SearchValue: request.ActivityGroupID,
		Limit:       100,
		Offset:      0,
	}

	todos, err := s.repo.ListTodos(ctx, params)

	if err != nil {
		return nil, utils.CustomError("failed to get todos", 400)
	}

	todosResp := make([]dtos.TodoResponse, len(todos))

	for i, todo := range todos {
		todoResp := dtos.ToTodoResponse(todo)
		todosResp[i] = todoResp
	}

	return todosResp, nil
}

func (s *todoServiceImpl) FindByID(ctx context.Context, id int32) (*dtos.TodoResponse, error) {
	todo, err := s.repo.GetTodo(ctx, id)

	if err != nil {
		return nil, utils.CustomError("failed to get todo", 400)
	}

	todoResp := dtos.ToTodoResponse(todo)

	return &todoResp, nil
}
