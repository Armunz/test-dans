package handler

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	repo "test-dans/internal/app/repository/jobs"
	mockRepo "test-dans/internal/app/repository/jobs/mock"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mockRepo.NewMockRepository(ctrl)

	type args struct {
		jobsRepo repo.Repository
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "When the function called, then it will return job usecase instance",
			args: args{
				jobsRepo: repoMock,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.args.jobsRepo)
			assert.NotNil(t, got)
		})
	}
}
