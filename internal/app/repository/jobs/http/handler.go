package http

import (
	"context"
	"strings"
	"time"

	"test-dans/model"
)

// GetJobDetail implements jobs.Repository
func (j *jobRepository) GetJobDetail(ctx context.Context, id string) (result model.Job, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(j.timeoutMs)*time.Millisecond)
	defer cancel()

	if err = j.safeguardDetail(ctx, id); err != nil {
		return
	}

	j.mu.RLock()
	defer j.mu.RUnlock()
	for _, job := range j.jobList {
		if id == job.ID {
			result = job
			break
		}
	}

	return
}

func (j *jobRepository) safeguardDetail(ctx context.Context, id string) error {
	select {
	case <-ctx.Done():
		return ErrCtxTimeout
	default:
	}

	if id == "" {
		return ErrEmptyID
	}

	return nil
}

// GetJobList implements jobs.Repository
func (j *jobRepository) GetJobList(ctx context.Context, page int, description string, location string, fullTime bool) (result model.JobList, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(j.timeoutMs)*time.Millisecond)
	defer cancel()

	if err = j.safeguardList(ctx); err != nil {
		return
	}

	j.mu.RLock()
	defer j.mu.RUnlock()

	var temp, candidate, res []model.Job

	if fullTime {
		for _, job := range j.jobList {
			if job.Type == "Full Time" {
				temp = append(temp, job)
			}
		}
	} else {
		temp = j.jobList
	}

	if description == "" && location == "" {
		candidate = temp
	} else {
		if description != "" {
			lowerDesc := strings.ToLower(description)
			for _, job := range temp {
				jobDescription := strings.ToLower(job.Description)
				index := strings.Index(jobDescription, lowerDesc)
				if index != -1 {
					candidate = append(candidate, job)
				}
			}
		}

		if location != "" {
			lowerLocation := strings.ToLower(location)
			for _, job := range temp {
				jobLocation := strings.ToLower(job.Location)
				if lowerLocation == jobLocation {
					candidate = append(candidate, job)
				}
			}
		}
	}

	candidate = removeDuplicateJob(candidate)
	totalPage := len(candidate)/PAGINATION + 1

	hasNext := true
	if page >= totalPage {
		hasNext = false
		result = model.JobList{
			TotalPage: totalPage,
			HasNext:   hasNext,
			Data:      nil,
		}
		return
	}

	if page == totalPage-1 {
		hasNext = false
	}

	startIndex := page * PAGINATION
	lastIndex := len(candidate) - 1
	counter := 0
	for {
		if counter == PAGINATION || startIndex > lastIndex {
			break
		}
		res = append(res, candidate[startIndex])
		startIndex++
		counter++
	}

	result = model.JobList{
		TotalPage: totalPage,
		HasNext:   hasNext,
		Data:      res,
	}

	return
}

func (j *jobRepository) safeguardList(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ErrCtxTimeout
	default:
	}

	return nil
}

func removeDuplicateJob(data []model.Job) (result []model.Job) {
	jobMap := make(map[string]struct{})

	for _, d := range data {
		if _, ok := jobMap[d.ID]; !ok {
			result = append(result, d)
			jobMap[d.ID] = struct{}{}
		}
	}

	return
}
