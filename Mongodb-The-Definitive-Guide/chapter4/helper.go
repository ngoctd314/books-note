package chapter4

import (
	"context"
	"log"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func getCollection(ctx context.Context) *mongo.Collection {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal("connect mongodb error:", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("ping mongodb error:", err)
	}

	collection := client.Database("testing").Collection("test")

	// drop before test
	collection.Drop(ctx)

	return collection

}

func breakLine() {
	log.Println(strings.Repeat("~", 40))
}
