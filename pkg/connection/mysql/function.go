package mysql

import (
	"database/sql"
	"test-dans/config"
)

func InitConnection(cfg config.MySQL) (database *sql.DB, err error) {
	dataSourceName := cfg.Username + `:` + cfg.Password + `@tcp(` + cfg.URL + `)/` + cfg.DBName
	database, err = sql.Open("mysql", dataSourceName)

	return
}
