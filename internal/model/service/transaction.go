package service

import (
	"context"

	mHandler "github.com/petchill/n-tier-rest-api-golang/internal/model/handler"
	mRes "github.com/petchill/n-tier-rest-api-golang/internal/model/response"
)

type TransactionService interface {
	Withdraw(ctx context.Context, payload mHandler.WithdrawReqBody) (mRes.WithdrawReply, error)
}
