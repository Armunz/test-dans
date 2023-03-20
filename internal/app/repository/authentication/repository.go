package authentication

import (
	"context"

	"test-dans/model"
)

type Repository interface {
	SetUser(ctx context.Context, username string, password string) (err error)
	GetUser(ctx context.Context, username string) (result model.UserLogin, err error)
}
