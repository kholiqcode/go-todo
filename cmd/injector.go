//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"

	"github.com/go-chi/chi/v5"
	v1 "github.com/kholiqcode/go-todolist/internal/activityGroup/delivery/http/v1"
	"github.com/kholiqcode/go-todolist/internal/app"
	"github.com/kholiqcode/go-todolist/utils"
)

func InitializeApp(route *chi.Mux, DB *sql.DB, config *utils.BaseConfig) (app.HttpServer, error) {
	panic(
		wire.Build(
			v1.ActivityGroupSet,
			app.AppSet,
		),
	)
}
