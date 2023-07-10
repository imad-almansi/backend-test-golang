package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/imad-almansi/backend-test-golang/pkg/handlers/results"
	"github.com/imad-almansi/backend-test-golang/pkg/model"
	"github.com/imad-almansi/backend-test-golang/pkg/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func HandleFacts(rw http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()

	filter := bson.D{}

	typeFilter := query.Get("type")
	if typeFilter != "" {
		filter = mongodb.FilterLiteral("type", typeFilter, filter)
	}
	foundFilter := query.Get("found")
	if foundFilter != "" {
		foundBoolean, err := strconv.ParseBool(foundFilter)
		if err != nil {
			results.Negative(rw, fmt.Errorf("found filter is an invalid boolean: %w", err), 400)
			return
		}
		filter = mongodb.FilterLiteral("found", foundBoolean, filter)
	}
	textFilter := query.Get("text")
	if textFilter != "" {
		filter = mongodb.FilterRegex("text", textFilter, filter)
	}

	findOptions := []*options.FindOptions{}

	limit := query.Get("limit")
	if limit != "" {
		limitInt, err := strconv.ParseInt(limit, 10, 64)
		if err != nil {
			results.Negative(rw, fmt.Errorf("limit is an invalid Int64: %w", err), 400)
			return
		}
		findOptions = append(findOptions, options.Find().SetLimit(limitInt))
	}

	cur, err := mongodb.Collection.Find(context.Background(), filter, findOptions...)
	if err != nil {
		results.Negative(rw, fmt.Errorf("Find failed: %w", err), 500)
		return
	}

	result := []model.Fact{}
	err = cur.All(context.Background(), &result)
	if err != nil {
		results.Negative(rw, fmt.Errorf("Decoding results failed: %w", err), 500)
		return
	}

	results.Positive(rw, result)
}
