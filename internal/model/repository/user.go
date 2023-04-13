package repository

import (
	"context"

	mUser "github.com/petchill/n-tier-rest-api-golang/internal/model/user"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int) (mUser.User, error)
}
