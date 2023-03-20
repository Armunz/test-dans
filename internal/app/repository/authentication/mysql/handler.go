package mysql

import (
	"context"
	"database/sql"
	"log"
	"time"

	"test-dans/model"
)

// GetUser implements login.Repository
func (m *mysqlUserRepo) GetUser(ctx context.Context, username string) (result model.UserLogin, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(m.timeoutMs)*time.Millisecond)
	defer cancel()

	if err = m.safeguardGet(ctx, username); err != nil {
		return
	}

	statement := `SELECT username, password FROM ` + m.tableName + ` WHERE username = ?`
	log.Println("[DEBUG] Statement: ", statement)
	err = m.conn.QueryRow(statement, username).Scan(&result.Username, &result.Password)

	// no need to return error when user not found
	if err == sql.ErrNoRows {
		err = nil
	}

	return
}

func (m *mysqlUserRepo) safeguardGet(ctx context.Context, username string) error {
	select {
	case <-ctx.Done():
		return ErrCtxTimeout
	default:
	}

	if m.conn == nil {
		return ErrDBConnNil
	}

	if m.tableName == "" {
		return ErrTableNameEmpty
	}

	if username == "" {
		return ErrUsernameEmpty
	}

	return nil
}

// SetUser implements authentication.Repository
func (m *mysqlUserRepo) SetUser(ctx context.Context, username string, password string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(m.timeoutMs)*time.Millisecond)
	defer cancel()

	if err = m.safeguardSet(ctx, username, password); err != nil {
		return
	}

	statement := `INSERT INTO ` + m.tableName + `(username, password) VALUES(?, ?)`
	res, err := m.conn.ExecContext(ctx, statement, username, password)
	if err != nil {
		return
	}

	_, err = res.RowsAffected()

	return
}

func (m *mysqlUserRepo) safeguardSet(ctx context.Context, username string, password string) error {
	select {
	case <-ctx.Done():
		return ErrCtxTimeout
	default:
	}

	if m.conn == nil {
		return ErrDBConnNil
	}

	if m.tableName == "" {
		return ErrTableNameEmpty
	}

	if username == "" {
		return ErrUsernameEmpty
	}

	if password == "" {
		return ErrEmptyPassword
	}

	return nil

}
