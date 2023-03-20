package handler

import (
	repo "test-dans/internal/app/repository/jobs"
	"test-dans/internal/app/usecase/jobs"
)

type jobsUsecase struct {
	jobsRepo repo.Repository
}

func New(jobsRepo repo.Repository) jobs.Usecase {
	return &jobsUsecase{
		jobsRepo: jobsRepo,
	}
}
