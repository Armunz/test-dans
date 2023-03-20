package http

import (
	"context"
	"encoding/json"
	"net/http"
	"test-dans/model"
	"time"

	"github.com/golang/gddo/httputil/header"
)

// Register implements delivery.Delivery
func (h *httpDelivery) Register(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(h.timeoutMs)*time.Millisecond)
	defer cancel()

	select {
	case <-ctx.Done():
		msg := "context timeout"
		http.Error(w, msg, http.StatusRequestTimeout)
		return
	default:
	}

	if r.Method != "POST" {
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

	var user model.UserLogin
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Username == "" || user.Password == "" {
		http.Error(w, "username or password is empty", http.StatusBadRequest)
		return
	}

	err = h.authUsecase.InsertUser(ctx, user.Username, user.Password)
	if err != nil {
		if err.Error() == "user already exist" {
			w.Write([]byte(err.Error()))
			return
		}
		
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
