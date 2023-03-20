package jobs

import (
	"context"

	"test-dans/model"
)

type Usecase interface {
	GetJobList(ctx context.Context, page int, description string, location string, fullTime bool) (result model.JobList, err error)
	GetJobDetail(ctx context.Context, id string) (result model.Job, err error)
}
