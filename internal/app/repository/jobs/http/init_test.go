package http

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	type args struct {
		url       string
		timeoutMS int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "When the function called, it will return job repository instance",
			args: args{
				url:       "http://localhost/api/recruitment/positions.json",
				timeoutMS: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.url, tt.args.timeoutMS)
			assert.NotNil(t, got)
		})
	}
}
