package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/imad-almansi/backend-test-golang/pkg/model"
	"github.com/imad-almansi/backend-test-golang/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

func HandleFacts(rw http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()

	filter := bson.D{}

	typeFilter := query.Get("type")
	if typeFilter != "" {
		fmt.Printf("%s\n", typeFilter)
	}

	cur, err := mongodb.Collection.Find(context.Background(), filter)
	if err != nil {
		log.Fatal("Find failed: ", err)
	}

	var results []model.Fact
	err = cur.All(context.Background(), &results)
	if err != nil {
		log.Fatal("Decode failed: ", err)
	}
	fmt.Println((results)[0])
}
