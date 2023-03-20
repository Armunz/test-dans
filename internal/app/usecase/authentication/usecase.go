package authentication

import (
	"context"
	"test-dans/model"
)

type Usecase interface {
	AuthenticateUser(ctx context.Context, username string, password string) (result model.UserLogin, err error)
	InsertUser(ctx context.Context, username string, password string) (err error)
}
