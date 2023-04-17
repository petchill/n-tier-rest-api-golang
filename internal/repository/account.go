package repository

import (
	"context"
	"errors"

	mAccount "github.com/petchill/n-tier-rest-api-golang/internal/model/account"
	mRepo "github.com/petchill/n-tier-rest-api-golang/internal/model/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type accountRepository struct {
	bankDB *mongo.Database
}

// GetAccountByID implements repository.AccountRepository
func (r accountRepository) GetAccountByID(ctx context.Context, id string) (mAccount.Account, error) {
	collection := r.bankDB.Collection("account")
	account := mAccount.Account{}

	query := bson.M{"id": id}
	err := collection.FindOne(ctx, query).Decode(&account)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return account, errors.New("Account is not found.")
		}
		return account, errors.New("Fail from account finding.")
	}
	return account, nil
}

// UpdateRemainAmountByID implements repository.AccountRepository
func (r accountRepository) UpdateRemainAmountByID(ctx context.Context, id string, remainAmount float64) error {
	collection := r.bankDB.Collection("account")

	query := bson.M{"id": id}
	payload := bson.M{
		"$set": bson.M{
			"remain_amount": remainAmount,
		},
	}
	err := collection.FindOneAndUpdate(ctx, query, payload).Err()

	if err != nil {
		return errors.New("Fail from account updating.")
	}
	return nil
}

func NewAccountRepository(bankDB *mongo.Database) mRepo.AccountRepository {
	return accountRepository{
		bankDB: bankDB,
	}
}
