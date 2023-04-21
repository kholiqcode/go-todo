package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/kholiqcode/go-todolist/pkg/logger"
)

func RunMigration(db *sql.DB, config *BaseConfig) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	LogAndPanicIfError(err, "cannot create new migrate instance")

	dbDriver := "mysql"

	m, err := migrate.NewWithDatabaseInstance(
		config.MIGRATION_URL,
		dbDriver, driver)

	LogAndPanicIfError(err, "failed to create mysql instance")

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		LogAndPanicIfError(err, "failed to run migrate up")
	}

	logger.LogInfo("db migrated successfully")
}
