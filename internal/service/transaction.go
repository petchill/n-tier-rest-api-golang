package service

import (
	"context"
	"errors"

	mAccount "github.com/petchill/n-tier-rest-api-golang/internal/model/account"
	mHandler "github.com/petchill/n-tier-rest-api-golang/internal/model/handler"
	mRepo "github.com/petchill/n-tier-rest-api-golang/internal/model/repository"
	mRes "github.com/petchill/n-tier-rest-api-golang/internal/model/response"
	mService "github.com/petchill/n-tier-rest-api-golang/internal/model/service"
	mUser "github.com/petchill/n-tier-rest-api-golang/internal/model/user"
)

type transactionService struct {
	accountRepo mRepo.AccountRepository
	userRepo    mRepo.UserRepository
}

// Withdraw implements service.TransactionService
func (s transactionService) Withdraw(ctx context.Context, payload mHandler.WithdrawReqBody) (mRes.WithdrawReply, error) {
	account, err := s.accountRepo.GetAccountByID(ctx, payload.AccountID)
	if err != nil {
		return mRes.WithdrawReply{}, err
	}
	user, err := s.userRepo.GetUserByID(ctx, account.UserID)
	if err != nil {
		return mRes.WithdrawReply{}, err
	}

	// validate withdrawer is match with account owner
	if userNameValid := validateUserName(user, payload.Withdrawer); !userNameValid {
		return mRes.WithdrawReply{}, errors.New("Withdrawers' name is not match with accounts' owner.")
	}

	// validate account available
	if accountValid := validateAccount(account); !accountValid {
		return mRes.WithdrawReply{}, errors.New("Account is invalid, Please contact banks' staff.")
	}

	// validate withdraw amount
	if withdrawAmountValid := validateWithdrawAmount(account, payload.Amount); !withdrawAmountValid {
		return mRes.WithdrawReply{}, errors.New("Money amount in your account is not enough for transaction.")
	}

	newRemainAmount := account.AmountRemain - payload.Amount

	// update account with new remain amount
	err = s.accountRepo.UpdateRemainAmountByID(ctx, account.ID, newRemainAmount)
	if err != nil {
		return mRes.WithdrawReply{}, err
	}

	reply := mRes.WithdrawReply{
		AccountID:      account.ID,
		UserName:       user.Name,
		WithdrawAmount: payload.Amount,
		RemainAmount:   newRemainAmount,
	}

	return reply, nil
}

func validateAccount(account mAccount.Account) bool {
	return account.AccountAvailable
}

func validateWithdrawAmount(account mAccount.Account, amount float64) bool {
	return account.AmountRemain >= amount
}

func validateUserName(user mUser.User, name string) bool {
	return user.Name == name
}

func NewTransactionService(accountRepo mRepo.AccountRepository, userRepo mRepo.UserRepository) mService.TransactionService {
	return transactionService{
		accountRepo: accountRepo,
		userRepo:    userRepo,
	}
}
