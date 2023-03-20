package http

import "test-dans/model"

type JobDetailResponse struct {
	Response model.Job `json:"response"`
}

type JobListResponse struct {
	Response model.JobList `json:"response"`
}
