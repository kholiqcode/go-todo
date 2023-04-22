//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"

	"github.com/go-chi/chi/v5"
	v1_activityGroup "github.com/kholiqcode/go-todolist/internal/activityGroup/delivery/http/v1"
	repo_activityGroup "github.com/kholiqcode/go-todolist/internal/activityGroup/repository"
	service_activityGroup "github.com/kholiqcode/go-todolist/internal/activityGroup/service"
	"github.com/kholiqcode/go-todolist/internal/app"
	v1_todo "github.com/kholiqcode/go-todolist/internal/todo/delivery/http/v1"
	repo_todo "github.com/kholiqcode/go-todolist/internal/todo/repository"
	service_todo "github.com/kholiqcode/go-todolist/internal/todo/service"
	"github.com/kholiqcode/go-todolist/utils"
)

func InitializeApp(route *chi.Mux, DB *sql.DB, config *utils.BaseConfig) (app.HttpServer, error) {
	panic(
		wire.Build(
			repo_activityGroup.ActivityGroupRepoSet,
			service_activityGroup.ActivityGroupServiceSet,
			v1_activityGroup.ActivityGroupHandlerSet,
			repo_todo.TodoRepoSet,
			service_todo.TodoServiceSet,
			v1_todo.TodoHandlerSet,
			app.AppSet,
		),
	)
}
