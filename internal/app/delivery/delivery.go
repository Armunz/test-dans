package delivery

import "net/http"

type Delivery interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	GetJobList(w http.ResponseWriter, r *http.Request)
	GetJobDetail(w http.ResponseWriter, r *http.Request)
}
