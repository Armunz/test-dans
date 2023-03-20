package handler

import (
	"context"
	"errors"
	"reflect"
	repo "test-dans/internal/app/repository/jobs"
	mockRepo "test-dans/internal/app/repository/jobs/mock"
	"test-dans/model"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_jobsUsecase_safeguardList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mockRepo.NewMockRepository(ctrl)

	timeoutCtx, cancel := context.WithCancel(context.Background())
	cancel()

	type fields struct {
		jobsRepo repo.Repository
	}
	type args struct {
		ctx  context.Context
		page int
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
				jobsRepo: repoMock,
			},
			args: args{
				ctx:  timeoutCtx,
				page: 0,
			},
			wantErr: true,
		},
		{
			name: "When repo is nil, then return error",
			fields: fields{
				jobsRepo: nil,
			},
			args: args{
				ctx:  context.Background(),
				page: 0,
			},
			wantErr: true,
		},
		{
			name: "When page is less than 0, then return error",
			fields: fields{
				jobsRepo: repoMock,
			},
			args: args{
				ctx:  context.Background(),
				page: -1,
			},
			wantErr: true,
		},
		{
			name: "When all is good, then return no error",
			fields: fields{
				jobsRepo: repoMock,
			},
			args: args{
				ctx:  context.Background(),
				page: 0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jobsUsecase{
				jobsRepo: tt.fields.jobsRepo,
			}
			if err := j.safeguardList(tt.args.ctx, tt.args.page); (err != nil) != tt.wantErr {
				t.Errorf("jobsUsecase.safeguardList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_jobsUsecase_safeguardDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mockRepo.NewMockRepository(ctrl)

	timeoutCtx, cancel := context.WithCancel(context.Background())
	cancel()

	type fields struct {
		jobsRepo repo.Repository
	}
	type args struct {
		ctx context.Context
		id  string
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
				jobsRepo: repoMock,
			},
			args: args{
				ctx: timeoutCtx,
				id:  "test",
			},
			wantErr: true,
		},
		{
			name: "When repo is nil, then return error",
			fields: fields{
				jobsRepo: nil,
			},
			args: args{
				ctx: context.Background(),
				id:  "test",
			},
			wantErr: true,
		},
		{
			name: "When id is empty, then return error",
			fields: fields{
				jobsRepo: repoMock,
			},
			args: args{
				ctx: context.Background(),
				id:  "",
			},
			wantErr: true,
		},
		{
			name: "When all is good, then return no error",
			fields: fields{
				jobsRepo: repoMock,
			},
			args: args{
				ctx: context.Background(),
				id:  "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jobsUsecase{
				jobsRepo: tt.fields.jobsRepo,
			}
			if err := j.safeguardDetail(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("jobsUsecase.safeguardDetail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_jobsUsecase_GetJobDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mockRepo.NewMockRepository(ctrl)

	type fields struct {
		jobsRepo repo.Repository
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		mockCalls  []func() *gomock.Call
		wantResult model.Job
		wantErr    bool
	}{
		{
			name:   "When safeguard fails, then return error",
			fields: fields{jobsRepo: nil},
			args: args{
				ctx: context.Background(),
				id:  "",
			},
			mockCalls:  []func() *gomock.Call{},
			wantResult: model.Job{},
			wantErr:    true,
		},
		{
			name: "When get job detail from repo got error, then return error",
			fields: fields{
				jobsRepo: repoMock,
			},
			args: args{
				ctx: context.Background(),
				id:  "test",
			},
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return repoMock.EXPECT().GetJobDetail(context.Background(), "test").Return(model.Job{}, errors.New("get job detail error"))
				},
			},
			wantResult: model.Job{},
			wantErr:    true,
		},
		{
			name: "When job detail with particular id is not found, then return empty",
			fields: fields{
				jobsRepo: repoMock,
			},
			args: args{
				ctx: context.Background(),
				id:  "test",
			},
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return repoMock.EXPECT().GetJobDetail(context.Background(), "test").Return(model.Job{}, nil)
				},
			},
			wantResult: model.Job{},
			wantErr:    false,
		},
		{
			name: "When job detail with particular id is found, then return job detail",
			fields: fields{
				jobsRepo: repoMock,
			},
			args: args{
				ctx: context.Background(),
				id:  "test",
			},
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return repoMock.EXPECT().GetJobDetail(context.Background(), "test").Return(model.Job{
						ID:          "test1",
						Type:        "Full Time",
						Location:    "Surabaya",
						Description: "lalala yeyeye Java lalala",
					}, nil)
				},
			},
			wantResult: model.Job{
				ID:          "test1",
				Type:        "Full Time",
				Location:    "Surabaya",
				Description: "lalala yeyeye Java lalala",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jobsUsecase{
				jobsRepo: tt.fields.jobsRepo,
			}

			if len(tt.mockCalls) > 0 {
				gomock.InOrder(func() []*gomock.Call {
					var mockCalls []*gomock.Call

					for _, call := range tt.mockCalls {
						mockCalls = append(mockCalls, call())
					}

					return mockCalls
				}()...)
			}

			gotResult, err := j.GetJobDetail(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("jobsUsecase.GetJobDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("jobsUsecase.GetJobDetail() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_jobsUsecase_GetJobList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mockRepo.NewMockRepository(ctrl)

	type fields struct {
		jobsRepo repo.Repository
	}
	type args struct {
		ctx         context.Context
		page        int
		description string
		location    string
		full_time   bool
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		mockCalls  []func() *gomock.Call
		wantResult model.JobList
		wantErr    bool
	}{
		{
			name: "When safeguard fails, then return error",
			fields: fields{
				jobsRepo: nil,
			},
			args: args{
				ctx:         context.Background(),
				page:        0,
				description: "",
				location:    "",
				full_time:   false,
			},
			wantResult: model.JobList{},
			wantErr:    true,
		},
		{
			name: "When get job list from repo got error, then return error",
			fields: fields{
				jobsRepo: repoMock,
			},
			args: args{
				ctx:         context.Background(),
				page:        0,
				description: "test",
				location:    "test",
				full_time:   false,
			},
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return repoMock.EXPECT().GetJobList(context.Background(), 0, "test", "test", false).Return(model.JobList{}, errors.New("get job list error"))
				},
			},
			wantResult: model.JobList{},
			wantErr:    true,
		},
		{
			name: "When get job list doesn't return any data, then return nil",
			fields: fields{
				jobsRepo: repoMock,
			},
			args: args{
				ctx:         context.Background(),
				page:        0,
				description: "test",
				location:    "test",
				full_time:   false,
			},
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return repoMock.EXPECT().GetJobList(context.Background(), 0, "test", "test", false).Return(model.JobList{
						TotalPage: 1,
						HasNext:   false,
						Data:      nil,
					}, nil)
				},
			},
			wantResult: model.JobList{
				TotalPage: 1,
				HasNext:   false,
				Data:      nil,
			},
			wantErr: false,
		},
		{
			name: "When all is good, then return job list",
			fields: fields{
				jobsRepo: repoMock,
			},
			args: args{
				ctx:         context.Background(),
				page:        0,
				description: "test",
				location:    "test",
				full_time:   false,
			},
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return repoMock.EXPECT().GetJobList(context.Background(), 0, "test", "test", false).Return(model.JobList{
						TotalPage: 1,
						HasNext:   false,
						Data: []model.Job{
							{
								ID:          "test1",
								Type:        "Full Time",
								Location:    "Surabaya",
								Description: "lalala yeyeye Java lalala",
							},
						},
					}, nil)
				},
			},
			wantResult: model.JobList{
				TotalPage: 1,
				HasNext:   false,
				Data: []model.Job{
					{
						ID:          "test1",
						Type:        "Full Time",
						Location:    "Surabaya",
						Description: "lalala yeyeye Java lalala",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jobsUsecase{
				jobsRepo: tt.fields.jobsRepo,
			}

			if len(tt.mockCalls) > 0 {
				gomock.InOrder(func() []*gomock.Call {
					var mockCalls []*gomock.Call

					for _, call := range tt.mockCalls {
						mockCalls = append(mockCalls, call())
					}

					return mockCalls
				}()...)
			}

			gotResult, err := j.GetJobList(tt.args.ctx, tt.args.page, tt.args.description, tt.args.location, tt.args.full_time)
			if (err != nil) != tt.wantErr {
				t.Errorf("jobsUsecase.GetJobList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("jobsUsecase.GetJobList() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
