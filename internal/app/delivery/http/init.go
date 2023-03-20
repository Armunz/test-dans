package http

import (
	"test-dans/internal/app/delivery"
	"test-dans/internal/app/usecase/authentication"
	"test-dans/internal/app/usecase/jobs"
)

type httpDelivery struct {
	authUsecase  authentication.Usecase
	jobsUsecase  jobs.Usecase
	jwtSecretKey string
	timeoutMs    int
}

func New(authUsecase authentication.Usecase, jobsUsecase jobs.Usecase, jwtSecretKey string, timeoutMs int) delivery.Delivery {
	return &httpDelivery{
		authUsecase:  authUsecase,
		jobsUsecase:  jobsUsecase,
		jwtSecretKey: jwtSecretKey,
		timeoutMs:    timeoutMs,
	}
}
