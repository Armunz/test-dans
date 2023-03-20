package http

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"test-dans/internal/app/repository/jobs"
	"test-dans/model"
)

type jobRepository struct {
	url       string
	mu        *sync.RWMutex
	jobList   []model.Job
	timeoutMs int
}

func New(url string, timeoutMS int) jobs.Repository {
	j := &jobRepository{
		url:       url,
		mu:        &sync.RWMutex{},
		jobList:   []model.Job{},
		timeoutMs: timeoutMS,
	}
	j.fillCache()

	return j
}

func (j *jobRepository) fillCache() {
	var err error
	var client = &http.Client{}
	var data []model.Job

	log.Println("Starting fill cache job list")
	request, err := http.NewRequest("GET", j.url, nil)
	if err != nil {
		log.Println("[error] failed to build request for get job list, ", err)
		return
	}

	response, err := client.Do(request)
	if err != nil {
		log.Println("[error] failed to get job list, ", err)
		return
	}
	defer response.Body.Close()

	log.Println("[DEBUG] Response Body: ", response.Body)

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		log.Println("[error] failed to parse job list, ", err)
		return
	}

	if len(data) == 0 {
		log.Println("[error] job list is empty")
		return
	}

	j.mu.Lock()
	j.jobList = data
	j.mu.Unlock()
}
