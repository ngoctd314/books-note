package chapter4

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Find is used perform queries in MongoDB. Querying returns a subset of documents in a collection
// Which documents get returned is determined by the first argument to find, which is a document specifying the query criteria.
func Find(ctx context.Context) {
	collection := getCollection(ctx)

	// An empty query document {} matches everything in the collection.
	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		log.Println(cur.Current)
	}

	// Insert a document with user name: admin
	filter = bson.D{{Key: "username", Value: "admin"}, {Key: "age", Value: 20}}
	_, err = collection.InsertOne(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("insert {username: admin}")

	// Find username admin
	filter = bson.D{{Key: "username", Value: "admin"}}
	cur, err = collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	breakLine()
	log.Println("find {username: admin}")
	for cur.Next(ctx) {
		log.Println(cur.Current)
	}

	// Find username admin1
	filter = bson.D{{Key: "username", Value: "admin1"}}
	cur, err = collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	breakLine()
	log.Println("find {username: admin1}")
	for cur.Next(ctx) {
		log.Println(cur.Current)
	}

	// Find username admin and age 20
	filter = bson.D{{Key: "username", Value: "admin"}, {Key: "age", Value: 20}}
	cur, err = collection.Find(ctx, filter)
	if err != nil {
		log.Println(err)
	}
	breakLine()
	log.Println("find {username: admin, age: 20}")
	for cur.Next(ctx) {
		log.Println(cur.Current)
	}
}

// Projection ...
// sometimes you don't need all of the key/value pairs in a document returned. In this case, you can pass a second argument
// to find (or findOne) specifying the keys you want.
func Projection(ctx context.Context) {
	collection := getCollection(ctx)

	// Insert a document
	_, err := collection.InsertOne(ctx, bson.D{{Key: "name", Value: "admin"}, {Key: "age", Value: 20}, {Key: "email", Value: "admin@gmail.com"}})
	if err != nil {
		log.Fatal(err)
	}

	// Find only name
	filter := bson.D{}
	opts := options.Find().SetProjection(bson.D{{Key: "name", Value: 1}, {Key: "age", Value: 1}, {Key: "_id", Value: 0}})
	cur, err := collection.Find(ctx, filter, opts)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		log.Println(cur.Current)
	}
}

// Limitations ...
// There are some restrictions on queries.
//
