package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	_handler "github.com/petchill/n-tier-rest-api-golang/internal/handler"
	_repo "github.com/petchill/n-tier-rest-api-golang/internal/repository"
	_service "github.com/petchill/n-tier-rest-api-golang/internal/service"
)

func main() {
	fmt.Println("hello this is first func")

	ctx := context.Background()

	// init mongo
	mongoOpts := options.Client().ApplyURI(os.Getenv("MONGO_CONN"))
	client, err := mongo.Connect(ctx, mongoOpts)
	if err != nil {
		log.Fatal("Cannot connect to mongodb client", err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal("Cannot disconnect to mongodb client", err)
		}
	}()

	bankDB := client.Database(os.Getenv("BANK_DB"))

	userRepo := _repo.NewUserRepository(bankDB)
	accountRepo := _repo.NewAccountRepository(bankDB)

	transactionService := _service.NewTransactionService(accountRepo, userRepo)

	transactionHanler := _handler.NewTransactionHandler(transactionService)

	http.HandleFunc("/withdraw", transactionHanler.WithdrawHandler)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
