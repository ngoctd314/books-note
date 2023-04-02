package chapter4

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// QueryCondition ...
// $gt, $lt, $gte, $lte, $ne are all comparison operators
func QueryCondition(ctx context.Context) {
	collection := getCollection(ctx)

	// Insert a document
	log.Println("insert {age: 20}")
	_, err := collection.InsertOne(ctx, bson.D{{Key: "age", Value: 20}})
	if err != nil {
		log.Fatal(err)
	}

	cur, err := collection.Find(ctx, bson.D{{Key: "age", Value: bson.D{{Key: "$gte", Value: 18}}}})
	if err != nil {
		log.Fatal(err)
	}
	breakLine()
	log.Println("find {age: {$gte: 18}}")
	for cur.Next(ctx) {
		log.Println(cur.Current)
	}

	cur, err = collection.Find(ctx, bson.D{{Key: "age", Value: bson.D{{Key: "$gte", Value: 22}}}})
	if err != nil {
		log.Fatal(err)
	}
	breakLine()
	log.Println("find {age: {$gte: 22}}")
	for cur.Next(ctx) {
		log.Println(cur.Current)
	}
}

// OrQuery ...
// $in can be used to query for a variety of values for a single key
// $nin opposite
// $or can be used to query for any of the given values across multiple keys
func OrQuery(ctx context.Context) {
	collection := getCollection(ctx)

	// Insert a document
	log.Println("insert {age: 20}, {age: 25}, {age: 30}")
	_, err := collection.InsertMany(ctx, []any{
		bson.D{{Key: "age", Value: 20}},
		bson.D{{Key: "age", Value: 25}},
		bson.D{{Key: "age", Value: 30}},
	})
	if err != nil {
		log.Fatal(err)
	}

	breakLine()
	log.Println("find {age: {$in: [18,19,20] }}")
	cur, err := collection.Find(ctx, bson.D{{Key: "age", Value: bson.D{{Key: "$in", Value: bson.A{18, 19, 20}}}}})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		log.Println(cur.Current)
	}
}

// NotQuery ...
// $not is a metaconditional: it can be applied on top of any other criteria.
func NotQuery(ctx context.Context) {
	collection := getCollection(ctx)

	// Insert a document
	_, err := collection.InsertOne(ctx, bson.M{"age": 23, "name": "ngoctd"})
	if err != nil {
		log.Fatal(err)
	}

	// Find a document
	var val1 any
	result := collection.FindOne(ctx, bson.D{{Key: "age", Value: 23}})
	result.Decode(&val1)
	log.Println(val1)

	// Find a document with $not operator
	result = collection.FindOne(ctx, bson.D{{Key: "age", Value: bson.D{{Key: "$gte", Value: 23}}}})
	var val2 any
	result.Decode(&val2)
	log.Println(val2)

	// Find a document with $not operator
	result = collection.FindOne(ctx, bson.D{{Key: "age", Value: bson.D{{Key: "$not", Value: bson.D{{Key: "$gte", Value: 23}}}}}})
	var val3 any
	result.Decode(&val3)
	log.Println(val3)
}
