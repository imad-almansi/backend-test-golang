package mongodb

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var Collection *mongo.Collection

// FilterLiteral creates a literal filter, if you want to append it to an existing filter (implicit and), pass the previous filter, otherwise pass nil or empty bson.D
func FilterLiteral(field string, value any, prevFilter bson.D) bson.D {
	if prevFilter == nil {
		prevFilter = bson.D{}
	}
	return append(prevFilter, bson.E{Key: field, Value: value})
}

// FilterRegex creates a regex evaluation filter, if you want to append it to an existing filter (implicit and), pass the previous filter, otherwise pass nil or empty bson.D
// value must be Perl compatible regular expressions
func FilterRegex(field string, value string, prevFilter bson.D) bson.D {
	if prevFilter == nil {
		prevFilter = bson.D{}
	}
	filter := bson.E{
		Key: field,
		Value: bson.D{
			{
				Key:   "$regex",
				Value: value,
			},
		},
	}

	return append(prevFilter, filter)
}
