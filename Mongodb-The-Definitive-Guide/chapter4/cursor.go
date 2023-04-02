package chapter4

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*
The database returns results from find using a cursor
*/

// LimitSkipSort ...
func LimitSkipSort(ctx context.Context) {
	collection := getCollection(ctx)

	collection.InsertMany(ctx, []any{
		bson.M{"no": 1},
		bson.M{"no": 2},
		bson.M{"no": 3},
		bson.M{"no": 4},
	})

	opts := options.Find().SetLimit(3).SetSkip(1).SetSort(bson.M{"no": -1})
	cur, err := collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		log.Println(cur.Current)
	}
}

/*
Using skip for a small number of documents is fine. But for a large number of results, skip can be slow
since it has to find and then discard all the skipped results.
MongDB does not yet support index with skips.
So large skip should be avoided. Often you can calculate the results of the next query based on the previous one.
*/
