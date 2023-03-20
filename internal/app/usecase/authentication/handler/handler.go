package handler

import (
	"context"
	"fmt"
	"log"

	passHelper "test-dans/internal/app/common/helper/password"
	"test-dans/model"
)

// AuthenticateUser implements login.Usecase
func (a *authUsecase) AuthenticateUser(ctx context.Context, username string, password string) (result model.UserLogin, err error) {
	if err = a.safeguard(ctx, username, password); err != nil {
		log.Println("[error] login safeguard failed, ", err)
		return
	}

	loginUser, err := a.authRepo.GetUser(ctx, username)
	if err != nil {
		log.Println("[error] failed to get user, ", err)
		return
	}

	// when user not found, return early with error
	if loginUser == (model.UserLogin{}) {
		err = fmt.Errorf("user doesn't exist")
		return
	}

	valid := passHelper.CheckPasswordHash(loginUser.Password, password)
	if !valid {
		err = fmt.Errorf("wrong password")
		return
	}

	result = loginUser

	return
}

// InsertUser implements authentication.Usecase
func (a *authUsecase) InsertUser(ctx context.Context, username string, password string) (err error) {
	if err = a.safeguard(ctx, username, password); err != nil {
		return
	}

	user, err := a.authRepo.GetUser(ctx, username)
	if err != nil {
		log.Println("[error] failed to get user, ", err)
		return
	}

	// when user already exist, then return early with error
	if user != (model.UserLogin{}) {
		err = fmt.Errorf("user already exist")
		return
	}

	hashPassword, err := passHelper.HashPassword(password)
	if err != nil {
		log.Println("[error] failed to generate hash password, ", err)
	}

	log.Println("[DEBUG] Hash Password: ", hashPassword)

	err = a.authRepo.SetUser(ctx, username, hashPassword)
	if err != nil {
		log.Println("[error] failed to insert user, ", err)
		return
	}

	return
}

func (l *authUsecase) safeguard(ctx context.Context, username string, password string) error {
	select {
	case <-ctx.Done():
		return ErrCtxTimeout
	default:
	}

	if l.authRepo == nil {
		return ErrRepoNil
	}

	if username == "" {
		return ErrEmptyUsername
	}

	if password == "" {
		return ErrEmptyPassword
	}

	return nil
}
