package http

import (
	"context"
	"encoding/json"
	"net/http"
	"test-dans/model"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/gddo/httputil/header"
)

// Login implements delivery.Delivery
func (h *httpDelivery) Login(w http.ResponseWriter, r *http.Request) {
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

	result, err := h.authUsecase.AuthenticateUser(ctx, user.Username, user.Password)
	if err != nil {
		if err.Error() == "user doesn't exist" {
			w.Write([]byte(err.Error()))
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &model.JWTClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
		Username: result.Username,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(h.jwtSecretKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    COOKIE_NAME,
		Value:   tokenString,
		Expires: expirationTime,
	})

	response := "Hello " + result.Username
	w.Write([]byte(response))
}
