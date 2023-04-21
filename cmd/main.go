package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/kholiqcode/go-todolist/utils"
)

func main() {
	r := chi.NewRouter()
	config := utils.CheckAndSetConfig("./", "app")
	DB := utils.ConnectDB(config)
	utils.RunMigration(DB, config)

	// ctx := context.Background()

	app, err := InitializeApp(r, DB, config)
	utils.LogAndPanicIfError(err, "failed when starting app")

	app.ListenAndServe()

}
