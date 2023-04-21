//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"

	"github.com/go-chi/chi/v5"
	v1_activityGroup "github.com/kholiqcode/go-todolist/internal/activityGroup/delivery/http/v1"
	service_activityGroup "github.com/kholiqcode/go-todolist/internal/activityGroup/service"
	repo_activityGroup "github.com/kholiqcode/go-todolist/internal/activityGroup/repository"
	"github.com/kholiqcode/go-todolist/internal/app"
	"github.com/kholiqcode/go-todolist/utils"
)

func InitializeApp(route *chi.Mux, DB *sql.DB, config *utils.BaseConfig) (app.HttpServer, error) {
	panic(
		wire.Build(
			repo_activityGroup.ActivityGroupRepoSet,
			service_activityGroup.ActivityGroupServiceSet,
			v1_activityGroup.ActivityGroupHandlerSet,
			app.AppSet,
		),
	)
}
