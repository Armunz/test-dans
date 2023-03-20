package mysql

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		conn      *sql.DB
		tableName string
		timeoutMs int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "When the function called, then it will return mysql repo instance",
			args: args{
				conn:      &sql.DB{},
				tableName: "test",
				timeoutMs: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.conn, tt.args.tableName, tt.args.timeoutMs)
			assert.NotNil(t, got)
		})
	}
}
