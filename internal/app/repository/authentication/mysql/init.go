package mysql

import (
	"database/sql"
	"test-dans/internal/app/repository/authentication"
)

type mysqlUserRepo struct {
	conn      *sql.DB
	tableName string
	timeoutMs int
}

func New(conn *sql.DB, tableName string, timeoutMs int) authentication.Repository {
	return &mysqlUserRepo{
		conn:      conn,
		tableName: tableName,
		timeoutMs: timeoutMs,
	}
}
