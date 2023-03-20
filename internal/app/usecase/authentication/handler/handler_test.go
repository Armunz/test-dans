package handler

import (
	"context"
	"errors"
	"reflect"
	repo "test-dans/internal/app/repository/authentication"
	mockRepo "test-dans/internal/app/repository/authentication/mock"
	"test-dans/model"
	"testing"

	"github.com/golang/mock/gomock"
)

func Test_loginUsecase_safeguard(t *testing.T) {
	timeoutCtx, cancel := context.WithCancel(context.Background())
	cancel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mockRepo.NewMockRepository(ctrl)

	type fields struct {
		authRepo repo.Repository
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
				authRepo: repoMock,
			},
			args: args{
				ctx:      timeoutCtx,
				username: "test",
				password: "test",
			},
			wantErr: true,
		},
		{
			name: "When repository is nil, then return error",
			fields: fields{
				authRepo: nil,
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
				authRepo: repoMock,
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
				authRepo: repoMock,
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
				authRepo: repoMock,
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
			l := &authUsecase{
				authRepo: tt.fields.authRepo,
			}
			if err := l.safeguard(tt.args.ctx, tt.args.username, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("loginUsecase.safeguard() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_loginUsecase_AuthenticateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mockRepo.NewMockRepository(ctrl)

	type fields struct {
		authRepo repo.Repository
	}
	type args struct {
		ctx      context.Context
		username string
		password string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		mockCalls  []func() *gomock.Call
		wantResult model.UserLogin
		wantErr    bool
	}{
		{
			name: "When safeguard fails, then return error",
			fields: fields{
				authRepo: repoMock,
			},
			args: args{
				ctx:      context.Background(),
				username: "",
				password: "",
			},
			wantResult: model.UserLogin{},
			wantErr:    true,
		},
		{
			name: "When get user from repo got error, then return error",
			fields: fields{
				authRepo: repoMock,
			},
			args: args{
				ctx:      context.Background(),
				username: "test",
				password: "test",
			},
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return repoMock.EXPECT().GetUser(context.Background(), "test").Return(model.UserLogin{}, errors.New("get user error"))
				},
			},
			wantResult: model.UserLogin{},
			wantErr:    true,
		},
		{
			name: "When user is not exist, then return error",
			fields: fields{
				authRepo: repoMock,
			},
			args: args{
				ctx:      context.Background(),
				username: "test",
				password: "test",
			},
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return repoMock.EXPECT().GetUser(context.Background(), "test").Return(model.UserLogin{}, nil)
				},
			},
			wantResult: model.UserLogin{},
			wantErr:    true,
		},
		{
			name: "When password is wrong, then return error",
			fields: fields{
				authRepo: repoMock,
			},
			args: args{
				ctx:      context.Background(),
				username: "test",
				password: "lalala",
			},
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return repoMock.EXPECT().GetUser(context.Background(), "test").Return(dummyUser, nil)
				},
			},
			wantResult: model.UserLogin{},
			wantErr:    true,
		},
		{
			name: "When all is good, then return login user",
			fields: fields{
				authRepo: repoMock,
			},
			args: args{
				ctx:      context.Background(),
				username: "test",
				password: "test",
			},
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return repoMock.EXPECT().GetUser(context.Background(), "test").Return(dummyUser, nil)
				},
			},
			wantResult: dummyUser,
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &authUsecase{
				authRepo: tt.fields.authRepo,
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

			gotResult, err := l.AuthenticateUser(tt.args.ctx, tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("loginUsecase.AuthenticateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("loginUsecase.AuthenticateUser() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_authUsecase_InsertUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mockRepo.NewMockRepository(ctrl)

	type fields struct {
		authRepo repo.Repository
	}
	type args struct {
		ctx      context.Context
		username string
		password string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		mockCalls []func() *gomock.Call
		wantErr   bool
	}{
		{
			name: "When safeguard fails, then return error",
			fields: fields{
				authRepo: repoMock,
			},
			args: args{
				ctx:      context.Background(),
				username: "",
				password: "",
			},
			wantErr: true,
		},
		{
			name: "When get user got error, then return error",
			fields: fields{
				authRepo: repoMock,
			},
			args: args{
				ctx:      context.Background(),
				username: "test",
				password: "test",
			},
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return repoMock.EXPECT().GetUser(context.Background(), "test").Return(model.UserLogin{}, errors.New("get user error"))
				},
			},
			wantErr: true,
		},
		{
			name: "When user already exist, then return early with error",
			fields: fields{
				authRepo: repoMock,
			},
			args: args{
				ctx:      context.Background(),
				username: "test",
				password: "test",
			},
			mockCalls: []func() *gomock.Call{
				func() *gomock.Call {
					return repoMock.EXPECT().GetUser(context.Background(), "test").Return(model.UserLogin{
						Username: "test",
						Password: "lalalayeyeye",
					}, nil)
				},
			},
			wantErr: true,
		},
		// because hashing always return different string, then can't be mocked
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &authUsecase{
				authRepo: tt.fields.authRepo,
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

			if err := a.InsertUser(tt.args.ctx, tt.args.username, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("authUsecase.InsertUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
