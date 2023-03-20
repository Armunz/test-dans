package model

import (
	"github.com/golang-jwt/jwt/v4"
)

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type JobList struct {
	TotalPage int
	HasNext   bool
	Data      []Job
}

type Job struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	URL         string `json:"url"`
	CreatedAt   string `json:"created_at"`
	Company     string `json:"company"`
	CompanyURL  string `json:"company_url"`
	Location    string `json:"location"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HowToApply  string `json:"how_to_apply"`
	CompanyLogo string `json:"company_logo"`
}

type JWTClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}
