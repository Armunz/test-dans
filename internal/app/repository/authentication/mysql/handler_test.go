package mysql

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"test-dans/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_mysqlUserRepo_safeguardGet(t *testing.T) {
	timeoutCtx, cancel := context.WithCancel(context.Background())
	cancel()

	type fields struct {
		conn      *sql.DB
		tableName string
		timeoutMs int
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "When context is timeout, then return error",
			fields: fields{
				conn:      &sql.DB{},
				tableName: "test",
				timeoutMs: 100,
			},
			args: args{
				ctx:      timeoutCtx,
				username: "test",
			},
			wantErr: true,
		},
		{
			name: "When sql connection is nil, then return error",
			fields: fields{
				conn:      nil,
				tableName: "test",
				timeoutMs: 100,
			},
			args: args{
				ctx:      context.Background(),
				username: "test",
			},
			wantErr: true,
		},
		{
			name: "When table name is empty, then return error",
			fields: fields{
				conn:      &sql.DB{},
				tableName: "",
				timeoutMs: 100,
			},
			args: args{
				ctx:      context.Background(),
				username: "test",
			},
			wantErr: true,
		},
		{
			name: "When username is empty, then return error",
			fields: fields{
				conn:      &sql.DB{},
				tableName: "test",
				timeoutMs: 100,
			},
			args: args{
				ctx:      context.Background(),
				username: "",
			},
			wantErr: true,
		},
		{
			name: "When all is good, then return no error",
			fields: fields{
				conn:      &sql.DB{},
				tableName: "test",
				timeoutMs: 100,
			},
			args: args{
				ctx:      context.Background(),
				username: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mysqlUserRepo{
				conn:      tt.fields.conn,
				tableName: tt.fields.tableName,
				timeoutMs: tt.fields.timeoutMs,
			}
			if err := m.safeguardGet(tt.args.ctx, tt.args.username); (err != nil) != tt.wantErr {
				t.Errorf("mysqlUserRepo.safeguard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlUserRepo_GetUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type fields struct {
		conn      *sql.DB
		tableName string
		timeoutMs int
	}
	type args struct {
		ctx      context.Context
		username string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		mockCall   func()
		wantResult model.UserLogin
		wantErr    bool
	}{
		{
			name: "When safeguard fails, then return error",
			fields: fields{
				conn:      db,
				tableName: "",
				timeoutMs: 100,
			},
			args: args{
				ctx:      context.Background(),
				username: "test",
			},
			mockCall:   func() {},
			wantResult: model.UserLogin{},
			wantErr:    true,
		},
		{
			name: "When get user from database return error, then return error",
			fields: fields{
				conn:      db,
				tableName: "test",
				timeoutMs: 100,
			},
			args: args{
				ctx:      context.Background(),
				username: "test",
			},
			mockCall: func() {
				mock.ExpectQuery("SELECT username, password FROM test WHERE username = ?").WithArgs("test").WillReturnError(errors.New("get user error"))
			},
			wantResult: model.UserLogin{},
			wantErr:    true,
		},
		{
			name: "When user is not found in database, then return empty data without error",
			fields: fields{
				conn:      db,
				tableName: "test",
				timeoutMs: 100,
			},
			args: args{
				ctx:      context.Background(),
				username: "test",
			},
			mockCall: func() {
				mock.ExpectQuery("SELECT username, password FROM test WHERE username = ?").WithArgs("test").WillReturnError(sql.ErrNoRows)
			},
			wantResult: model.UserLogin{},
			wantErr:    false,
		},
		{
			name: "When all is good, then return username and password that match in database",
			fields: fields{
				conn:      db,
				tableName: "test",
				timeoutMs: 100,
			},
			args: args{
				ctx:      context.Background(),
				username: "test",
			},
			mockCall: func() {
				rows := sqlmock.NewRows([]string{"username", "password"}).AddRow("test", "lalala")
				mock.ExpectQuery("SELECT username, password FROM test WHERE username = ?").WithArgs("test").WillReturnRows(rows)
			},
			wantResult: model.UserLogin{
				Username: "test",
				Password: "lalala",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mysqlUserRepo{
				conn:      tt.fields.conn,
				tableName: tt.fields.tableName,
				timeoutMs: tt.fields.timeoutMs,
			}

			tt.mockCall()
			gotResult, err := m.GetUser(tt.args.ctx, tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlUserRepo.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("mysqlUserRepo.GetUser() = %v, want %v", gotResult, tt.wantResult)
			}
			// we make sure that all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func Test_mysqlUserRepo_safeguardSet(t *testing.T) {
	timeoutCtx, cancel := context.WithCancel(context.Background())
	cancel()

	type fields struct {
		conn      *sql.DB
		tableName string
		timeoutMs int
	}
	type args struct {
		ctx      context.Context
		username string
		password string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "When context is timeout, then return error",
			fields: fields{
				conn:      &sql.DB{},
				tableName: "test",
				timeoutMs: 100,
			},
			args: args{
				ctx:      timeoutCtx,
				username: "test",
				password: "test",
			},
			wantErr: true,
		},
		{
			name: "When sql connection is nil, then return error",
			fields: fields{
				conn:      nil,
				tableName: "test",
				timeoutMs: 100,
			},
			args: args{
				ctx:      context.Background(),
				username: "test",
				password: "test",
			},
			wantErr: true,
		},
		{
			name: "When table name is empty, then return error",
			fields: fields{
				conn:      &sql.DB{},
				tableName: "",
				timeoutMs: 100,
			},
			args: args{
				ctx:      context.Background(),
				username: "test",
				password: "test",
			},
			wantErr: true,
		},
		{
			name: "When username is empty, then return error",
			fields: fields{
				conn:      &sql.DB{},
				tableName: "test",
				timeoutMs: 100,
			},
			args: args{
				ctx:      context.Background(),
				username: "",
				password: "test",
			},
			wantErr: true,
		},
		{
			name: "When password is empty, then return error",
			fields: fields{
				conn:      &sql.DB{},
				tableName: "test",
				timeoutMs: 100,
			},
			args: args{
				ctx:      context.Background(),
				username: "test",
				password: "",
			},
			wantErr: true,
		},
		{
			name: "When all is good, then return no error",
			fields: fields{
				conn:      &sql.DB{},
				tableName: "test",
				timeoutMs: 100,
			},
			args: args{
				ctx:      context.Background(),
				username: "test",
				password: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mysqlUserRepo{
				conn:      tt.fields.conn,
				tableName: tt.fields.tableName,
				timeoutMs: tt.fields.timeoutMs,
			}
			if err := m.safeguardSet(tt.args.ctx, tt.args.username, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("mysqlUserRepo.safeguardSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_mysqlUserRepo_SetUser(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type fields struct {
		conn      *sql.DB
		tableName string
		timeoutMs int
	}
	type args struct {
		ctx      context.Context
		username string
		password string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		mockCall func()
		wantErr  bool
	}{
		{
			name: "When safeguard fails, then return error",
			fields: fields{
				conn:      &sql.DB{},
				tableName: "",
				timeoutMs: 100,
			},
			args: args{
				ctx:      context.Background(),
				username: "",
				password: "",
			},
			mockCall: func() {},
			wantErr:  true,
		},
		{
			name: "When set user to database got error, then return error",
			fields: fields{
				conn:      db,
				tableName: "test",
				timeoutMs: 100,
			},
			args: args{
				ctx:      context.Background(),
				username: "test",
				password: "test",
			},
			mockCall: func() {
				statement := `INSERT INTO test(username, password) VALUES(?, ?)`
				mock.ExpectExec(statement).WithArgs("test", "test").WillReturnError(errors.New("set user error"))
			},
			wantErr: true,
		},
		{
			name: "When all is good, then set user to database and return no error",
			fields: fields{
				conn:      db,
				tableName: "test",
				timeoutMs: 100,
			},
			args: args{
				ctx:      context.Background(),
				username: "test",
				password: "test",
			},
			mockCall: func() {
				statement := `INSERT INTO test(username, password) VALUES(?, ?)`
				mock.ExpectExec(statement).WithArgs("test", "test").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &mysqlUserRepo{
				conn:      tt.fields.conn,
				tableName: tt.fields.tableName,
				timeoutMs: tt.fields.timeoutMs,
			}
			tt.mockCall()
			if err := m.SetUser(tt.args.ctx, tt.args.username, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("mysqlUserRepo.SetUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			// we make sure that all expectations were met
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("there were unfulfilled expectations: %s", err)
			}
		})
	}
}
