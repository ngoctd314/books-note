package chapter5

import (
	"context"
	"log"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Indexes ...
func Indexes(ctx context.Context) {
	collection := getCollection(ctx)
	rand.Seed(time.Now().UnixNano())

	// for i := 0; i < 1000000; i++ {
	// 	_, err := collection.InsertOne(ctx, bson.M{
	// 		"i":        i,
	// 		"username": fmt.Sprintf("user-%d", i),
	// 		"age":      20 + rand.Intn(10),
	// 		"created":  primitive.NewDateTimeFromTime(time.Now()),
	// 	})
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// 	if i%100000 == 0 {
	// 		log.Println("i: ", i)
	// 	}
	// }

	/*
		Creating an Index
		indexes have their price: write operations (inserts, updates and deletes) that modify an indexes field will take longer.
		This is because in addition to updating the document, MongoDB has to update indexes when you data changes.
		Typocally, the tradeoff is worth it. The tricky part becomes figuring out which fields to index.

		To choose which fields to create indexes for, look through your frequent queries and queries that need to be fast and try to find a common set of keys from those.
	*/
	createIndexResult, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.M{"username": 1},
	})
	if err != nil {
		log.Fatal("create index: ", err)
	}
	log.Println("create index:", createIndexResult)

	//    The purpose of an index is to make your queries as efficient as possible. For many query patterns it is necessary to build
	//    indexes based on two or more keys. For example, an index keeps all of its values in a sorted order, so it makes sorting
	//    documents by the indexed key much faster.
	//    However, an index can only help with sorting if it is a prefix of the sort. For example, the index on "username" wouldn't help
	//    much for this sort:
	//    db.users.find().sort({"age": 1, "username": 1})

	//    This sorts by "age" and then "username", so a strict sorting by "username" isn't terribly helpful.
	//    To optimize this sort, you could make an index on "age" and "username"
	//    db.users.createIndex({"age": 1, "username": 1})

	//    This is called a compound index and is useful if your query has multiple sort directions or multiple keys in the criteria.

}
