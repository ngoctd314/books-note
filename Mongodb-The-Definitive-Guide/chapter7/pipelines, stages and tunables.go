package chapter7

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

/*
MongoDB provides powerful support for running analytics natively using the aggregation framework.
*/

// Aggregate ...
func Aggregate(ctx context.Context) {
	collection := getCollection(ctx)
	insertMany(ctx, collection)

	cur, err := collection.Aggregate(ctx, []any{
		bson.M{"$match": bson.M{"state": "a"}},
		bson.M{"$group": bson.M{"_id": "$city", "totalAge": bson.M{"$count": struct{}{}}}},
		bson.M{"$limit": 1},
	})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		log.Println(cur.Current)
	}
}

func insertMany(ctx context.Context, collection *mongo.Collection) {
	docs := []any{}
	listState := []string{"a", "b", "c", "d", "e", "f"}
	cities := []string{"HN", "HCM", "DN"}
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 1000; i++ {
		docs = append(docs, bson.M{
			"age":   rand.Intn(100) + 1,
			"name":  fmt.Sprintf("name-%d", i),
			"log":   fmt.Sprintf("loc-%d", i),
			"state": listState[rand.Intn(len(listState))],
			"city":  cities[rand.Intn(len(cities))],
		})
	}
	collection.InsertMany(ctx, docs)
}
