package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/joho/godotenv"
	_handler "github.com/petchill/n-tier-rest-api-golang/internal/handler"
	_repo "github.com/petchill/n-tier-rest-api-golang/internal/repository"
	_service "github.com/petchill/n-tier-rest-api-golang/internal/service"
)

func main() {
	ctx := context.Background()

	godotenv.Load()

	// init mongoกิ
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

	// init repository
	userRepo := _repo.NewUserRepository(bankDB)
	accountRepo := _repo.NewAccountRepository(bankDB)

	// init service
	transactionService := _service.NewTransactionService(accountRepo, userRepo)

	// init handler
	transactionHanler := _handler.NewTransactionHandler(transactionService)

	http.HandleFunc("/withdraw", transactionHanler.WithdrawHandler)
	fmt.Println("server start on port 5000 ...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
