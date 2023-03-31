package chapter3

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Document ...
type Document struct {
	Title string
}

// InsertOne will add an "_id" key to the document (if you don't supply one) and store the document in MongoDB
func (d Document) InsertOne(ctx context.Context) {
	// get collection
	collection := mongoDBConnection(ctx)
	result, err := collection.InsertOne(ctx, bson.D{{"title", d.Title}})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("insert result:", result)
}

// InsertMany enables you to pass an array of documents to the database
// Make more efficient to reduce the db round trip
func (d Document) InsertMany(ctx context.Context) {
	// get collection
	collection := mongoDBConnection(ctx)
	n := 10
	data := make([]any, 0, n)
	for i := 0; i < n; i++ {
		data = append(data, bson.D{{"title", fmt.Sprintf("%s-%d", d.Title, i)}, {"_id", "1"}})
	}
	result, err := collection.InsertMany(ctx, data, &options.InsertManyOptions{
		Ordered: new(bool),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("insert result:", result)
}
