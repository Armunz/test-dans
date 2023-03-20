package http

import (
	"context"
	"reflect"
	"sync"
	"test-dans/model"
	"testing"
)

func Test_jobRepository_safeguardDetail(t *testing.T) {
	timeoutCtx, cancel := context.WithCancel(context.Background())
	cancel()

	type fields struct {
		url       string
		mu        *sync.RWMutex
		jobList   []model.Job
		timeoutMs int
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
			name: "When context timeout, then return error",
			fields: fields{
				url:       "http://localhost:9999",
				mu:        &sync.RWMutex{},
				jobList:   []model.Job{},
				timeoutMs: 100,
			},
			args: args{
				ctx: timeoutCtx,
				id:  "test",
			},
			wantErr: true,
		},
		{
			name: "When ID is empty, then return error",
			fields: fields{
				url:       "http://localhost:9999",
				mu:        &sync.RWMutex{},
				jobList:   []model.Job{},
				timeoutMs: 100,
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
				url:       "http://localhost:9999",
				mu:        &sync.RWMutex{},
				jobList:   []model.Job{},
				timeoutMs: 100,
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
			j := &jobRepository{
				url:       tt.fields.url,
				mu:        tt.fields.mu,
				jobList:   tt.fields.jobList,
				timeoutMs: tt.fields.timeoutMs,
			}
			if err := j.safeguardDetail(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("jobRepository.safeguardDetail() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_jobRepository_safeguardList(t *testing.T) {
	timeoutCtx, cancel := context.WithCancel(context.Background())
	cancel()

	type fields struct {
		url       string
		mu        *sync.RWMutex
		jobList   []model.Job
		timeoutMs int
	}
	type args struct {
		ctx context.Context
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
				url:       "http://localhost:9999",
				mu:        &sync.RWMutex{},
				jobList:   []model.Job{},
				timeoutMs: 100,
			},
			args: args{
				ctx: timeoutCtx,
			},
			wantErr: true,
		},
		{
			name: "When all is good, then return no error",
			fields: fields{
				url:       "http://localhost:9999",
				mu:        &sync.RWMutex{},
				jobList:   []model.Job{},
				timeoutMs: 100,
			},
			args: args{
				ctx: context.Background(),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jobRepository{
				url:       tt.fields.url,
				mu:        tt.fields.mu,
				jobList:   tt.fields.jobList,
				timeoutMs: tt.fields.timeoutMs,
			}
			if err := j.safeguardList(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("jobRepository.safeguardList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_jobRepository_GetJobDetail(t *testing.T) {
	timeoutCtx, cancel := context.WithCancel(context.Background())
	cancel()

	type fields struct {
		url       string
		mu        *sync.RWMutex
		jobList   []model.Job
		timeoutMs int
	}
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult model.Job
		wantErr    bool
	}{
		{
			name: "When safeguard fails, then return error",
			fields: fields{
				url:       "http://localhost:9999",
				mu:        &sync.RWMutex{},
				jobList:   dummyJobList,
				timeoutMs: 100,
			},
			args: args{
				ctx: timeoutCtx,
				id:  "test2",
			},
			wantResult: model.Job{},
			wantErr:    true,
		},
		{
			name: "When id is found in job list, then return the job",
			fields: fields{
				url:       "http://localhost:9999",
				mu:        &sync.RWMutex{},
				jobList:   dummyJobList,
				timeoutMs: 100,
			},
			args: args{
				ctx: context.Background(),
				id:  "test5",
			},
			wantResult: model.Job{
				ID:          "test5",
				Type:        "Full Time",
				Location:    "Surabaya",
				Description: "hahaha huhu hahaha java hehe",
			},
			wantErr: false,
		},
		{
			name: "When id is not found in job list, then return empty",
			fields: fields{
				url:       "http://localhost:9999",
				mu:        &sync.RWMutex{},
				jobList:   dummyJobList,
				timeoutMs: 100,
			},
			args: args{
				ctx: context.Background(),
				id:  "test99",
			},
			wantResult: model.Job{},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jobRepository{
				url:       tt.fields.url,
				mu:        tt.fields.mu,
				jobList:   tt.fields.jobList,
				timeoutMs: tt.fields.timeoutMs,
			}
			gotResult, err := j.GetJobDetail(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("jobRepository.GetJobDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("jobRepository.GetJobDetail() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_jobRepository_GetJobList(t *testing.T) {
	timeoutCtx, cancel := context.WithCancel(context.Background())
	cancel()

	type fields struct {
		url       string
		mu        *sync.RWMutex
		jobList   []model.Job
		timeoutMs int
	}
	type args struct {
		ctx         context.Context
		page        int
		description string
		location    string
		fullTime    bool
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult model.JobList
		wantErr    bool
	}{
		{
			name: "When safeguard fails, then return error",
			fields: fields{
				url:       "http://localhost:9999",
				mu:        &sync.RWMutex{},
				jobList:   dummyJobList,
				timeoutMs: 100,
			},
			args: args{
				ctx:         timeoutCtx,
				page:        0,
				description: "java",
				location:    "surabaya",
				fullTime:    false,
			},
			wantResult: model.JobList{},
			wantErr:    true,
		},
		{
			name: "When full time is false and all other argument is empty, then return all job list on particular pagination",
			fields: fields{
				url:       "http://localhost:9999",
				mu:        &sync.RWMutex{},
				jobList:   dummyJobList,
				timeoutMs: 100,
			},
			args: args{
				ctx:         context.Background(),
				page:        0,
				description: "",
				location:    "",
				fullTime:    false,
			},
			wantResult: model.JobList{
				TotalPage: 1,
				HasNext:   false,
				Data:      dummyJobList,
			},
			wantErr: false,
		},
		{
			name: "When full time is true and other argument is empty, then return all job that type is full time on particular pagination",
			fields: fields{
				url:       "http://localhost:9999",
				mu:        &sync.RWMutex{},
				jobList:   dummyJobList,
				timeoutMs: 100,
			},
			args: args{
				ctx:         context.Background(),
				page:        0,
				description: "",
				location:    "",
				fullTime:    true,
			},
			wantResult: model.JobList{
				TotalPage: 1,
				HasNext:   false,
				Data:      wantResultFullTimeOnly,
			},
			wantErr: false,
		},
		{
			name: "When pagination exceed total page, then return nil data",
			fields: fields{
				url:       "http://localhost:9999",
				mu:        &sync.RWMutex{},
				jobList:   dummyJobList,
				timeoutMs: 100,
			},
			args: args{
				ctx:         context.Background(),
				page:        1,
				description: "",
				location:    "",
				fullTime:    false,
			},
			wantResult: model.JobList{
				TotalPage: 1,
				HasNext:   false,
				Data:      nil,
			},
			wantErr: false,
		},
		{
			name: "When full time is true and description is java, then return match data",
			fields: fields{
				url:       "http://localhost:9999",
				mu:        &sync.RWMutex{},
				jobList:   dummyJobList,
				timeoutMs: 100,
			},
			args: args{
				ctx:         context.Background(),
				page:        0,
				description: "java",
				location:    "",
				fullTime:    true,
			},
			wantResult: model.JobList{
				TotalPage: 1,
				HasNext:   false,
				Data:      wantResultFullTimeJava,
			},
			wantErr: false,
		},
		{
			name: "When full time is true and description is java and location is surabaya, then return match data",
			fields: fields{
				url:       "http://localhost:9999",
				mu:        &sync.RWMutex{},
				jobList:   dummyJobList,
				timeoutMs: 100,
			},
			args: args{
				ctx:         context.Background(),
				page:        0,
				description: "java",
				location:    "surabaya",
				fullTime:    true,
			},
			wantResult: model.JobList{
				TotalPage: 1,
				HasNext:   false,
				Data:      wantResultFullTimeJavaSurabaya,
			},
			wantErr: false,
		},
		{
			name: "When full time is false and location is surabaya, then return match data",
			fields: fields{
				url:       "http://localhost:9999",
				mu:        &sync.RWMutex{},
				jobList:   dummyJobList,
				timeoutMs: 100,
			},
			args: args{
				ctx:         context.Background(),
				page:        0,
				description: "",
				location:    "surabaya",
				fullTime:    false,
			},
			wantResult: model.JobList{
				TotalPage: 1,
				HasNext:   false,
				Data:      wantResultFullTimeFalseSurabaya,
			},
			wantErr: false,
		},
		{
			name: "When full time is false and description is react and location is Jakarta, then return match data",
			fields: fields{
				url:       "http://localhost:9999",
				mu:        &sync.RWMutex{},
				jobList:   dummyJobList,
				timeoutMs: 100,
			},
			args: args{
				ctx:         context.Background(),
				page:        0,
				description: "react",
				location:    "jakarta",
				fullTime:    false,
			},
			wantResult: model.JobList{
				TotalPage: 1,
				HasNext:   false,
				Data:      wantResultJakartaReact,
			},
			wantErr: false,
		},
		{
			name: "When arguments didn't match with any data, then return nil",
			fields: fields{
				url:       "http://localhost:9999",
				mu:        &sync.RWMutex{},
				jobList:   dummyJobList,
				timeoutMs: 100,
			},
			args: args{
				ctx:         context.Background(),
				page:        0,
				description: "assembly",
				location:    "madura",
				fullTime:    false,
			},
			wantResult: model.JobList{
				TotalPage: 1,
				HasNext:   false,
				Data:      nil,
			},
			wantErr: false,
		},
		// This test case i used when pagination is 3
		// {
		// 	name: "When pagination is on last page, then return match data",
		// 	fields: fields{
		// 		url:       "http://localhost:9999",
		// 		mu:        &sync.RWMutex{},
		// 		jobList:   dummyJobList,
		// 		timeoutMs: 100,
		// 	},
		// 	args: args{
		// 		ctx:         context.Background(),
		// 		page:        2,
		// 		description: "",
		// 		location:    "",
		// 		fullTime:    false,
		// 	},
		// 	wantResult: model.JobList{
		// 		TotalPage: 3,
		// 		HasNext:   false,
		// 		Data:      wantResultLastPage,
		// 	},
		// 	wantErr: false,
		// },
		// {
		// 	name: "When pagination is on middle page, then return match data",
		// 	fields: fields{
		// 		url:       "http://localhost:9999",
		// 		mu:        &sync.RWMutex{},
		// 		jobList:   dummyJobList,
		// 		timeoutMs: 100,
		// 	},
		// 	args: args{
		// 		ctx:         context.Background(),
		// 		page:        1,
		// 		description: "",
		// 		location:    "",
		// 		fullTime:    false,
		// 	},
		// 	wantResult: model.JobList{
		// 		TotalPage: 3,
		// 		HasNext:   true,
		// 		Data:      wantResultMiddlePage,
		// 	},
		// 	wantErr: false,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &jobRepository{
				url:       tt.fields.url,
				mu:        tt.fields.mu,
				jobList:   tt.fields.jobList,
				timeoutMs: tt.fields.timeoutMs,
			}
			gotResult, err := j.GetJobList(tt.args.ctx, tt.args.page, tt.args.description, tt.args.location, tt.args.fullTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("jobRepository.GetJobList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("jobRepository.GetJobList() = %v, \nwant %v", gotResult, tt.wantResult)
			}
		})
	}
}
