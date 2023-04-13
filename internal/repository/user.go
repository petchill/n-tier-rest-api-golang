package repository

import (
	"context"
	"errors"

	mRepo "github.com/petchill/n-tier-rest-api-golang/internal/model/repository"
	mUser "github.com/petchill/n-tier-rest-api-golang/internal/model/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	bankDB *mongo.Database
}

// GetUserByID implements user.Repository
func (r userRepository) GetUserByID(ctx context.Context, id int) (mUser.User, error) {
	collection := r.bankDB.Collection("user")
	user := mUser.User{}

	query := bson.M{"id": id}
	err := collection.FindOne(ctx, query).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, errors.New("ไม่พบผู้ใช้ในระบบ")
		}
		return user, errors.New("การค้นหาผู้ใช้ผิดพลาด")
	}
	return user, nil
}

func NewUserRepository(bankDB *mongo.Database) mRepo.UserRepository {
	return userRepository{
		bankDB: bankDB,
	}
}
