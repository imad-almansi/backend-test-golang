package mongodb

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestFilterLiteral(t *testing.T) {
	cases := map[string]struct {
		field      string
		value      any
		prevFilter bson.D
		expected   bson.D
	}{
		"filter a string field": {
			field:    "type",
			value:    "abc",
			expected: bson.D{{Key: "type", Value: "abc"}},
		},
		"filter a boolean field": {
			field:    "found",
			value:    true,
			expected: bson.D{{Key: "found", Value: true}},
		},
		"filter 2 string fields": {
			field:      "type",
			value:      "xyz",
			prevFilter: bson.D{{Key: "type", Value: "abc"}},
			expected:   bson.D{{Key: "type", Value: "abc"}, {Key: "type", Value: "xyz"}},
		},
		"filter a string field and a boolean field": {
			field:      "found",
			value:      true,
			prevFilter: bson.D{{Key: "type", Value: "abc"}},
			expected:   bson.D{{Key: "type", Value: "abc"}, {Key: "found", Value: true}},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result := FilterLiteral(c.field, c.value, c.prevFilter)

			assert.Equal(t, c.expected, result)
		})
	}
}

func TestFilterRegex(t *testing.T) {
	cases := map[string]struct {
		field      string
		value      string
		prevFilter bson.D
		expected   bson.D
	}{
		"filter a string field with wildcard regex": {
			field: "text",
			value: ".*",
			expected: bson.D{{
				Key: "text",
				Value: bson.D{
					{
						Key:   "$regex",
						Value: ".*",
					},
				},
			}},
		},
		"filter a string field with a regex": {
			field: "text",
			value: "hello",
			expected: bson.D{{
				Key: "text",
				Value: bson.D{
					{
						Key:   "$regex",
						Value: "hello",
					},
				},
			}},
		},
		"filter a string field with multipl regex filters": {
			field: "text",
			value: "world",
			prevFilter: bson.D{{
				Key: "text",
				Value: bson.D{
					{
						Key:   "$regex",
						Value: "hello",
					},
				},
			}},
			expected: bson.D{
				{
					Key: "text",
					Value: bson.D{
						{
							Key:   "$regex",
							Value: "hello",
						},
					},
				},
				{
					Key: "text",
					Value: bson.D{
						{
							Key:   "$regex",
							Value: "world",
						},
					},
				},
			},
		},
		"filter with 2 fields with regex": {
			field: "type",
			value: "abc",
			prevFilter: bson.D{{
				Key: "text",
				Value: bson.D{
					{
						Key:   "$regex",
						Value: "hello",
					},
				},
			}},
			expected: bson.D{
				{
					Key: "text",
					Value: bson.D{
						{
							Key:   "$regex",
							Value: "hello",
						},
					},
				},
				{
					Key: "type",
					Value: bson.D{
						{
							Key:   "$regex",
							Value: "abc",
						},
					},
				},
			},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			result := FilterRegex(c.field, c.value, c.prevFilter)

			assert.Equal(t, c.expected, result)
		})
	}
}
