package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(cfg *BaseConfig) *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	dbDriver := "mysql"
	dbc, err := sql.Open(dbDriver, dsn)
	LogAndPanicIfError(err, "failed when connecting to database")

	err = dbc.Ping()
	LogAndPanicIfError(err, "failed when ping to database")

	return dbc
}
