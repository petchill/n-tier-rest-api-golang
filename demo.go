package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	mUser "github.com/petchill/n-tier-rest-api-golang/internal/model/user"
)

func main2() {

	ctx := context.Background()

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

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		collection := bankDB.Collection("user")
		users := []mUser.User{}

		cursor, err := collection.Find(ctx, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = cursor.All(ctx, &users)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
		return
	})

}
