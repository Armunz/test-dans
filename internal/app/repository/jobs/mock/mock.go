// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_jobs is a generated GoMock package.
package mock_jobs

import (
	context "context"
	reflect "reflect"
	model "test-dans/model"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// GetJobDetail mocks base method.
func (m *MockRepository) GetJobDetail(ctx context.Context, id string) (model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobDetail", ctx, id)
	ret0, _ := ret[0].(model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobDetail indicates an expected call of GetJobDetail.
func (mr *MockRepositoryMockRecorder) GetJobDetail(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobDetail", reflect.TypeOf((*MockRepository)(nil).GetJobDetail), ctx, id)
}

// GetJobList mocks base method.
func (m *MockRepository) GetJobList(ctx context.Context, page int, description, location string, full_time bool) (model.JobList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobList", ctx, page, description, location, full_time)
	ret0, _ := ret[0].(model.JobList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobList indicates an expected call of GetJobList.
func (mr *MockRepositoryMockRecorder) GetJobList(ctx, page, description, location, full_time interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobList", reflect.TypeOf((*MockRepository)(nil).GetJobList), ctx, page, description, location, full_time)
}