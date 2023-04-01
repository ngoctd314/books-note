package modeling

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
we made use of  bson.A, bson.D and bson.M which represents arrays, documents and maps
*/

// Podcast ...
type Podcast struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   []string           `bson:"tags:omitempty"`
}

// Episode ...
type Episode struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Poscast     primitive.ObjectID `bson:"podcast,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Description string             `bson:"title,omitempty"`
	Duration    int32              `bson:"duration,omitempty"`
}

// Fn ...
func Fn() {
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	database := client.Database("quickstart")
	postcastCollection := database.Collection("podcasts")
	episodesCollection := database.Collection("episodes")
	episodesCollection.Drop(ctx)
	postcastCollection.Drop(ctx)

	// var episodes []Episode
	// cur, err := episodesCollection.Find(ctx, bson.M{"duration": bson.D{{Key: "$gt", Value: 25}}})
	// if err != nil {
	// 	panic(err)
	// }
	// if err = cur.All(ctx, &episodes); err != nil {
	// 	panic(err)
	// }
	// log.Println(episodes)

	podcast := Podcast{
		Title:  "The Polyglot Developer",
		Author: "Nic Raboy",
		Tags:   []string{"development", "programming", "coding"},
	}
	insertResult, err := postcastCollection.InsertOne(ctx, podcast)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(insertResult.InsertedID)
}
