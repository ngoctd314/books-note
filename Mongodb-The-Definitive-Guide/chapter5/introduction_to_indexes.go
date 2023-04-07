package chapter5

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

// Indexes ...
func Indexes(ctx context.Context) {
	collection := getCollection(ctx)
	rand.Seed(time.Now().UnixNano())

	// wg := sync.WaitGroup{}
	// wg.Add(6)
	// for i := 0; i <= 9; i++ {
	// 	go func(i int) {
	// 		defer wg.Done()
	// 		for j := 100000 * i; j < (i+1)*100000; j++ {
	// 			_, err := collection.InsertOne(ctx, bson.M{
	// 				"age": 1 + rand.Intn(100),
	// 			})
	// 			if err != nil {
	// 				log.Println(err)
	// 			}
	// 			if j%100000 == 0 {
	// 				log.Println("j: ", j)
	// 			}
	// 		}
	// 	}(i)
	// }
	// wg.Wait()

	// Creating an Index
	// indexes have their price: write operations (inserts, updates and deletes) that modify an indexes field will take longer.
	// An indexes can make a dramatic difference in query times.However, indexes have their price: write operations (inserts, updates and deletes) that modify an indexed field will take longer
	// This is because in addition to updating the document, MongoDB has to update indexes when your data change
	// Typically, the tradeoff is worth it. The tricky part becomes figuring out which fields to index.

	// To choose which fields to create indexes for, look through your frequent queries and queries that need to be fast and try to find a common set of keys from those.

	// without index
	// find: 2.47ms		find and sort: 4.2s

	// with index
	// find: 2.9ms	    find and sort: 3.8ms
	// collection.Indexes().DropAll(ctx)
	// cur, _ := collection.Indexes().List(ctx)
	// for cur.Next(ctx) {
	// 	log.Println(cur.Current)
	// }

	// collection.Indexes().DropAll(ctx)
	// query with and without indexes
	now := time.Now()
	// opts := options.Find().SetSort(bson.M{"age": 1})
	cnt := 0
	cur, err := collection.Find(ctx, bson.M{"age": bson.M{"$gt": 20}})
	// cur, err := collection.Find(ctx, bson.M{"age": 21})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("time query:", time.Since(now)) // 21.5s, 13.7ms
	for cur.Next(ctx) {
		// log.Println(cur.Current)
		cnt++
	}
	fmt.Println(cnt)

	// log.Println("time query:", time.Since(now)) // 21.5s, 13.7ms
	// multi-key map passed in for ordered parameter keys

	// collection.Indexes().DropAll(ctx)
	// index return document in sorted order
	// createIndexResult, err := collection.Indexes().CreateOne(ctx, mongo.IndexModel{
	// 	Keys: bson.D{{Key: "age", Value: 1}},
	// })
	// if err != nil {
	// 	log.Fatal("create index: ", err)
	// }
	// log.Println("create index:", createIndexResult)

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

/*
How MongoDB selects an Index
Now, let's take a look at how MongoDB chooses an index to satisfy a query. When a query comes in, MongDB looks at the query's shape.
The shape has to do with what fields are being searched on and additional information, such as whether or not there is a sort. Based on that information
the system identifies a set of candidate indexes that it might be able to use in satisfying the query.

Let's assume we have a query come in, and three of our five indexes are identified as candidates for this query. MongoDB will then create three query plans,
one for each of these indexes, and run the query in three parallel threads, each using different index. The objective here is to see which one is able return results the fastest.

The server maintains a of query plans. A winning plan is stored in the cache for future use for queries of that shape.

## Using Compound Indexes
Compound indexes are a little more complicated to think about than single-key indexes, but they are very powerful.
*/
