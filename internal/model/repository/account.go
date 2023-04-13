package repository

import (
	"context"

	mAccount "github.com/petchill/n-tier-rest-api-golang/internal/model/account"
)

type AccountRepository interface {
	GetAccountByID(ctx context.Context, id string) (mAccount.Account, error)
	UpdateRemainAmountByID(ctx context.Context, id string, remainAmount float64) error
}
