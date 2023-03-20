package jobs

import (
	"context"
	"test-dans/model"
)

type Repository interface {
	GetJobList(ctx context.Context, page int, description string, location string, full_time bool) (result model.JobList, err error)
	GetJobDetail(ctx context.Context, id string) (result model.Job, err error)
}
