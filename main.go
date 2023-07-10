package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/imad-almansi/backend-test-golang/pkg/handlers"
	"github.com/imad-almansi/backend-test-golang/pkg/mongodb"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	if mongodb.Collection != nil {
		return
	}

	ctx := context.Background()
	user := os.Getenv("MONGO_READ_USER")
	password := os.Getenv("MONGO_READ_PASS")
	if user == "" || password == "" {
		log.Fatal("database user(MONGO_READ_USER) or password(MONGO_READ_PASS) is not defined")
	}
	database := os.Getenv("DB_NAME")
	collection := os.Getenv("DB_COLLECTION")
	if database == "" {
		log.Fatal("database name(DB_NAME) is not defined")
	}
	if collection == "" {
		log.Fatal("database collection(DB_COLLECTION) is not defined")
	}
	host := os.Getenv("DB_HOST")
	if host == "" {
		log.Printf("database host(DB_HOST) is not defined, using localhost:27017 as default")
		host = "localhost:27017"
	}

	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s/", host)).SetAuth(options.Credential{
		Username: user,
		Password: password,
	})
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("client is unable to connect to client", err)
	}

	mongodb.Collection = client.Database(database).Collection(collection)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/facts", handlers.HandleFacts)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	log.Println("starting server on port 8080")
	log.Fatal(srv.ListenAndServe())
}
