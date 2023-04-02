package chapter4

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Null ...
func Null(ctx context.Context) {}

// RegularExpression ...
func RegularExpression(ctx context.Context) {}

// QueryingArrays ...
// Querying for elements of an array is designed to behave the way querying for scalars does.
func QueryingArrays(ctx context.Context) {
	collection := getCollection(ctx)

	// Insert
	collection.InsertOne(ctx, bson.M{"fruit": []any{"apple", "banana", "peach"}})

	// Find
	result := collection.FindOne(ctx, bson.M{"fruit": "apple"})
	var val any
	result.Decode(&val)
	log.Println(val)
}

// QueryingArraysAllOperation ...
// If you need to match arrays by more than one element, you can use $all
// This allows you to match a list of elements.
func QueryingArraysAllOperation(ctx context.Context) {
	collection := getCollection(ctx)

	collection.InsertMany(ctx, []any{
		bson.M{"fruit": []any{"apple", "banana", "peach"}},
		bson.M{"fruit": []any{"apple", "kumquat", "orange"}},
		bson.M{"fruit": []any{"cherry", "banana", "apple"}},
	})

	// We can find all documents with both "apple" and "banana" elements
	cur, err := collection.Find(ctx, bson.M{"fruit": bson.M{"$all": []any{"apple", "banana"}}})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		log.Println(cur.Current)
	}

	breakLine()
	// We can also query by exact match using the entire array. However, exact match will not match a document
	cur, err = collection.Find(ctx, bson.M{"fruit": []any{"apple", "banana"}})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		log.Println(cur.Current)
	}

	breakLine()
	// We can also query by exact match using the entire array. However, exact match will not match a document
	cur, err = collection.Find(ctx, bson.M{"fruit": []any{"apple", "banana", "peach"}})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		log.Println(cur.Current)
	}
}

// QueryingArraysSizeOperator ...
// A useful conditional for querying arrays is $size, which allows you to query for arrays of a given size
func QueryingArraysSizeOperator(ctx context.Context) {
	collection := getCollection(ctx)

	collection.InsertMany(ctx, []any{
		bson.M{"fruit": []any{"apple", "banana", "peach"}},
		bson.M{"fruit": []any{"apple", "kumquat", "orange"}},
		bson.M{"fruit": []any{"cherry", "banana", "apple"}},
	})

	cur, err := collection.Find(ctx, bson.M{"fruit": bson.M{"$size": 3}})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		log.Println(cur.Current)
	}

	breakLine()
	cur, err = collection.Find(ctx, bson.M{"fruit": bson.M{"$size": 1}})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		log.Println(cur.Current)
	}
}

// QueryingArraysSliceOperator ...
// $slice operator can be used to return a subset of elements for an array key
func QueryingArraysSliceOperator(ctx context.Context) {
	collection := getCollection(ctx)

	_, err := collection.InsertMany(ctx, []any{
		bson.M{"fruit": []any{"apple", "banana", "peach"}},
		bson.M{"fruit": []any{"apple", "kumquat", "orange"}},
		bson.M{"fruit": []any{"cherry", "banana", "apple"}},
	})
	if err != nil {
		log.Println(err)
	}

	// we wanted the first 2 fruits
	opts := options.FindOne().SetProjection(bson.M{"fruit": bson.M{"$slice": 2}})
	rs := collection.FindOne(ctx, bson.D{{Key: "fruit", Value: "apple"}}, opts)
	var val any
	rs.Decode(&val)
	log.Println(val)

	// we wanted the last 2 fruits
	opts = options.FindOne().SetProjection(bson.M{"fruit": bson.M{"$slice": -2}})
	rs = collection.FindOne(ctx, bson.D{{Key: "fruit", Value: "apple"}}, opts)
	var val1 any
	rs.Decode(&val1)
	log.Println(val1)

	// we wanted the middle of the results by tanking an offset and the number of elements to return
	opts = options.FindOne().SetProjection(bson.M{"fruit": bson.M{"$slice": []any{1, 2}}}) // skip 1 element and return 2 element
	rs = collection.FindOne(ctx, bson.D{{Key: "fruit", Value: "apple"}}, opts)
	var val2 any
	rs.Decode(&val2)
	log.Println(val2)
}

// ArrayAndRangeQuery ...
func ArrayAndRangeQuery(ctx context.Context) {}

// QueryingOnEmbedded ...
func QueryingOnEmbedded(ctx context.Context) {
	collection := getCollection(ctx)

	_, err := collection.InsertMany(ctx, []any{
		bson.M{"name": bson.M{"first": "Joe", "last": "Schmoe"}, "age": 23},
	})

	if err != nil {
		log.Fatal(err)
	}

	// Query by full subdocument
	cur, err := collection.Find(ctx, bson.M{"name": bson.M{"first": "Joe", "last": "Schmoe"}})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		log.Println(cur.Current)
	}

	// Query by embedded key
	// query documents can contain dots. Which mean "reach into an embedded document".
	cur, err = collection.Find(ctx, bson.M{"name.first": "Joe"})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		log.Println(cur.Current)
	}
}

// QueryingArraysEmbedded ...
func QueryingArraysEmbedded(ctx context.Context) {
	collection := getCollection(ctx)

	_, err := collection.InsertMany(ctx, []any{
		bson.M{"comments": []any{bson.M{"author": "Joe", "score": 3}, bson.M{"author": "Mary", "score": 6}}, "content": "content 1"},
	})
	if err != nil {
		log.Fatal(err)
	}

	// query comment store > 5
	opts := options.Find().SetProjection(bson.M{"comments": 1, "_id": 0})
	cur, err := collection.Find(ctx, bson.M{"comments": bson.M{"$elemMatch": bson.M{"score": bson.M{"$gte": 6}}}}, opts)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		log.Println(cur.Current)
	}
}
