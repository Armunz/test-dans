package http

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"test-dans/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/gddo/httputil/header"
	"github.com/gorilla/mux"
)

// GetJobDetail implements delivery.Delivery
func (h *httpDelivery) GetJobDetail(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.timeoutMs)*time.Millisecond)
	defer cancel()

	select {
	case <-ctx.Done():
		msg := "context timeout"
		http.Error(w, msg, http.StatusRequestTimeout)
		return
	default:
	}

	if r.Method != "GET" {
		msg := "Unsupported http method"
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	vars := mux.Vars(r)
	jobID := vars["id"]

	c, err := r.Cookie(COOKIE_NAME)
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tknStr := c.Value
	claims := &model.JWTClaims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return h.jwtSecretKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	result, err := h.jobsUsecase.GetJobDetail(ctx, jobID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := JobDetailResponse{
		Response: result,
	}

	jsonData, err := json.Marshal(res)
	if err != nil {
		log.Println("[error] failed to do json marshal of job detail, ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)

}

// GetJobList implements delivery.Delivery
func (h *httpDelivery) GetJobList(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.timeoutMs)*time.Millisecond)
	defer cancel()

	select {
	case <-ctx.Done():
		msg := "context timeout"
		http.Error(w, msg, http.StatusRequestTimeout)
		return
	default:
	}

	if r.Method != "GET" {
		msg := "Unsupported http method"
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	if r.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(r.Header, "Content-Type")
		if value != "application/json" {
			msg := "Content-Type header is not application/json"
			http.Error(w, msg, http.StatusUnsupportedMediaType)
			return
		}
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "page is required", http.StatusBadRequest)
		return
	}

	description := r.URL.Query().Get("description")
	location := r.URL.Query().Get("location")
	fullTime, err := strconv.ParseBool(r.URL.Query().Get("full_time"))
	if err != nil {
		http.Error(w, "full_time should be boolean", http.StatusBadRequest)
		return
	}

	c, err := r.Cookie(COOKIE_NAME)
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tknStr := c.Value
	claims := &model.JWTClaims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return h.jwtSecretKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	result, err := h.jobsUsecase.GetJobList(ctx, page, description, location, fullTime)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := JobListResponse{
		Response: result,
	}

	jsonData, err := json.Marshal(res)
	if err != nil {
		log.Println("[error] failed to do json marshal of job list, ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonData)

}
