package handler

import (
	"testing"

	repo "test-dans/internal/app/repository/authentication"
	mockRepo "test-dans/internal/app/repository/authentication/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mockRepo.NewMockRepository(ctrl)

	type args struct {
		loginRepo repo.Repository
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "When the function called, then it will return login usecase instance",
			args: args{
				loginRepo: repoMock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.loginRepo)
			assert.NotNil(t, got)
		})
	}
}
