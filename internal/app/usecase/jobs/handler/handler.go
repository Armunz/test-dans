package handler

import (
	"context"
	"log"
	"test-dans/model"
)

// GetJobDetail implements jobs.Usecase
func (j *jobsUsecase) GetJobDetail(ctx context.Context, id string) (result model.Job, err error) {
	if err = j.safeguardDetail(ctx, id); err != nil {
		return
	}

	result, err = j.jobsRepo.GetJobDetail(ctx, id)
	if err != nil {
		log.Println("[error] failed to get job detail, ", err)
	}

	return
}

// GetJobList implements jobs.Usecase
func (j *jobsUsecase) GetJobList(ctx context.Context, page int, description string, location string, full_time bool) (result model.JobList, err error) {
	if err = j.safeguardList(ctx, page); err != nil {
		return
	}

	result, err = j.jobsRepo.GetJobList(ctx, page, description, location, full_time)
	if err != nil {
		log.Println("[error] failed to get job list, ", err)
	}

	return
}

func (j *jobsUsecase) safeguardDetail(ctx context.Context, id string) error {
	select {
	case <-ctx.Done():
		return ErrCtxTimeout
	default:
	}

	if j.jobsRepo == nil {
		return ErrRepoNil
	}

	if id == "" {
		return ErrEmptyID
	}

	return nil
}

func (j *jobsUsecase) safeguardList(ctx context.Context, page int) error {
	select {
	case <-ctx.Done():
		return ErrCtxTimeout
	default:
	}

	if j.jobsRepo == nil {
		return ErrRepoNil
	}

	if page < 0 {
		return ErrPageInvalid
	}

	return nil
}
