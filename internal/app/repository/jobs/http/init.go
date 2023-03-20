package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	var data []model.Job

	log.Println("Starting fill cache job list")
	// make GET request
	response, err := http.Get(j.url)
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("[error] failed to read response body, ", err)
	}

	buf := bytes.NewBuffer(body)
	err = json.NewDecoder(buf).Decode(&data)
	if err != nil {
		log.Println("[error] failed to parse job list, ", err)
		return
	}
	response.Body.Close()

	if len(data) == 0 {
		log.Println("[error] job list is empty")
		return
	}

	j.mu.Lock()
	j.jobList = data
	j.mu.Unlock()
}
