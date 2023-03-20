package handler

import (
	repo "test-dans/internal/app/repository/authentication"
	"test-dans/internal/app/usecase/authentication"
)

type authUsecase struct {
	authRepo repo.Repository
}

func New(authRepo repo.Repository) authentication.Usecase {
	return &authUsecase{
		authRepo: authRepo,
	}
}
