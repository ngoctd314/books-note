package chapter3

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Document ...
type Document struct {
	Title string
	Count int
}

// InsertOne will add an "_id" key to the document (if you don't supply one) and store the document in MongoDB
// MongoDB does minimal checks on data being inserted: it checks the document's basic structure and adds an "_id" field if one doesn't exist.
// All documents must be smaller than 16MB.
func (d Document) InsertOne(ctx context.Context) {
	// get collection
	collection := mongoDBConnection(ctx)
	result, err := collection.InsertOne(ctx, bson.D{{Key: "title", Value: d.Title}, {Key: "count", Value: d.Count}})
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
		data = append(data, bson.D{{Key: "title", Value: fmt.Sprintf("%s-%d", d.Title, i)}, {Key: "count", Value: d.Count}})
	}
	result, err := collection.InsertMany(ctx, data, &options.InsertManyOptions{
		Ordered: new(bool),
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("insert result:", result)
}

// InsertManyUnOrdered when performing a bulk insert using insertMany, if a document halfway through the array produces an error
// of some type, what happens on whether you have opted for ordered or unordered operations.
// Specify true for the key "ordered" to ensure documents are inserted in the order they are provided.
// Specify false and MongoDB may reorder the inserts to increase performance
// Ordered inserts is the default. If a document produces an insertion error, no documents beyond that point in the array will be inserted.
// For unordered inserts, MongoDB will attempt to insert all documents, regardless of whether some insertions produce errors.
func (d Document) InsertManyUnOrdered(ctx context.Context, ordered bool) {
	conn := mongoDBConnection(ctx)
	n := 10
	data := make([]any, 0, n)
	for i := 0; i < n; i++ {
		data = append(data, bson.D{{Key: "title", Value: fmt.Sprintf("%s-%d", d.Title, i)}, {Key: "_id", Value: "0"}}) // duplicate key
	}
	result, err := conn.InsertMany(ctx, data, &options.InsertManyOptions{
		Ordered: &ordered, // default true
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("insert results:", result)

}

// Find ...
func (d Document) Find(ctx context.Context) {
	conn := mongoDBConnection(ctx)
	cur, err := conn.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		log.Println("results:", cur.Current)
	}
}

// FindByID ...
func (d Document) FindByID(ctx context.Context, id string) {
	conn := mongoDBConnection(ctx)
	objID, _ := primitive.ObjectIDFromHex(id)
	cur, err := conn.Find(ctx, bson.D{{Key: "_id", Value: objID}})
	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(ctx) {
		log.Println("results:", cur.Current)
	}
}

// DeleteOne will delete the first document found that matches the filter.
// Which document is found first depends on several factors, including the order in which the documents were inserted
// what updates were made to the documents, and what indexes are specified.
func (d Document) DeleteOne(ctx context.Context, id any) {
	conn := mongoDBConnection(ctx)
	rs, err := conn.DeleteOne(ctx, bson.D{{Key: "_id", Value: id}})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("delete one result:", rs.DeletedCount)

}

// DeleteMany delete all the documents that match a filter
func (d Document) DeleteMany(ctx context.Context, title any) {
	conn := mongoDBConnection(ctx)

	filter := bson.D{}
	filter = append(filter, primitive.E{
		Key:   "title",
		Value: title,
	})

	rs, err := conn.DeleteMany(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("delete many result:", rs.DeletedCount)
}

// Drop it is possible to use deleteMany to remove all documents in a collection
// However, if you want to clear an entire collection, it is faster to drop it
// Once data has been removed, it is gone forever. There is no way to undo a delete or drop operation or recover deleted documents (except backup)
func (d Document) Drop(ctx context.Context) {
	conn := mongoDBConnection(ctx)

	err := conn.Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}
}

// UpdateOne take a filter document as their first parameter and a modifier document
// When to use: only certain portions of a document need to be updated. You can update specific fields in a documents using atomic update operations.
// Updating a document is atomic: if two updates happen at the same time, whichever one reaches the server first will be applied, and then the next will be applied
// The last update will "win"
func (d Document) UpdateOne(ctx context.Context, id string) {
	conn := mongoDBConnection(ctx)

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "count", Value: 1}}}}
	rs, err := conn.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("update result:", rs)
}

// UpdateMany take a filter document as their first parameter and a modifier document
func (d Document) UpdateMany(ctx context.Context, id string) {
	conn := mongoDBConnection(ctx)

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "count", Value: 1}}}, {Key: "$set", Value: bson.D{{Key: "title", Value: "updated"}}}}
	rs, err := conn.UpdateMany(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("update result:", rs)
}

// func (d Document) ReplaceOne(ctx context.Context) {}

// SetOperator sets the value of a field. If the field does not yet exist, it will be created. This can be handy for updating schemas or adding user-defined keys.
func (d Document) SetOperator(ctx context.Context, id, newTitle string) {
	conn := mongoDBConnection(ctx)

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.D{}
	// $inc is similar to $set, but it is designed for incrementing (and decrementing) numbers.
	// $inc can be used only on values of type integer, long double, or decimal
	// Also, the value of the $inc key must be a number
	update = append(update, primitive.E{Key: "$inc", Value: bson.D{{Key: "count", Value: 1}}})
	update = append(update, primitive.E{Key: "$set", Value: bson.D{{Key: "title", Value: newTitle}}})
	update = append(update, primitive.E{Key: "$set", Value: bson.D{{Key: "name", Value: bson.D{{Key: "email", Value: "ngoctd@gmail.com"}, {Key: "address", Value: "Thanh Xuan, Ha Noi"}}}}})

	rs, err := conn.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("update result:", rs)
}

// UnSetOperator $unset delete matching field
func (d Document) UnSetOperator(ctx context.Context, id string, field string) {
	conn := mongoDBConnection(ctx)

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.D{{Key: "$unset", Value: bson.D{{Key: field}}}}
	rs, err := conn.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("update result:", rs)
}

// SetEmbeddedDocument you can also use $set to reach in and change embedded documents
func (d Document) SetEmbeddedDocument(ctx context.Context, id string) {
	conn := mongoDBConnection(ctx)

	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name.email", Value: "ngoc@gmail.com"}}}}
	rs, err := conn.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("update result:", rs)
}

// PushOperator $push adds elements to the end of an array if the array exists and creates a new array if it does not.
func (d Document) PushOperator(ctx context.Context) {
	conn := mongoDBConnection(ctx)

	data := bson.D{}
	data = append(data, primitive.E{Key: "title", Value: "A blog post"})
	data = append(data, primitive.E{Key: "content", Value: "Learn mongodb array"})

	// insert
	rs, err := conn.InsertOne(ctx, data)
	if err != nil {
		log.Println("insert one error:", err)
	}

	// find
	post, err := conn.Find(ctx, bson.D{{Key: "_id", Value: rs.InsertedID}})
	if err != nil {
		log.Println("insert one error:", err)
	}
	for post.Next(ctx) {
		log.Println(post.Current)
	}

	filter := bson.D{{Key: "_id", Value: rs.InsertedID}}
	update := bson.D{{Key: "$push", Value: bson.D{{Key: "comments", Value: bson.D{{Key: "name", Value: "ngoctd"}, {Key: "email", Value: "ngoctd@gmail.com"}, {Key: "content", Value: "nice post."}}}}}}
	_, err = conn.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("insert one error:", err)
	}

	post, err = conn.Find(ctx, bson.D{{Key: "_id", Value: rs.InsertedID}})
	if err != nil {
		log.Println("insert one error:", err)
	}
	for post.Next(ctx) {
		log.Println(post.Current)
	}

	update = bson.D{{Key: "$push", Value: bson.D{{Key: "hourly", Value: bson.D{{Key: "$each", Value: bson.A{1, 2, 3, 4}}}}}}}
	_, err = conn.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Println("insert one error:", err)
	}
	post, err = conn.Find(ctx, bson.D{{Key: "_id", Value: rs.InsertedID}})
	if err != nil {
		log.Println("insert one error:", err)
	}
	for post.Next(ctx) {
		log.Println(post.Current)
	}
}

// Upsert is a special type of update.
// If no document is found that matches filter, a new document will be created by combining the criteria and updated documents
// If a matching document is found, it will be updated normally.
// Upserts can be handy because they can eliminate the need to "seed" your collection: you can often have the same code create and update documents.
// Without Upserts, we need making a round trip to the database, plus sending an update or insert. If we are running this code in multiple processes
// we are also subject to a race condition where more than one document can be inserted for a given URL.
// Upsert is atomic.
func (d Document) Upsert(ctx context.Context) {
	conn := mongoDBConnection(ctx)
	filter := bson.D{{Key: "url", Value: "/blog-1"}}
	update := bson.D{{Key: "$inc", Value: bson.D{{Key: "pageviews", Value: 1}}}}
	update = append(update, primitive.E{
		// SetOnInsertOperator sometimes a field needs to be set when a document is created, but not changed on subsequent updates
		// $setOnInsert is an operator that only sets the value of a field when the document is being inserted.
		Key: "$setOnInsert",
		Value: bson.D{{
			Key:   "createAt",
			Value: primitive.NewDateTimeFromTime(time.Now()),
		}},
	})

	upsert := true
	rs, err := conn.UpdateOne(ctx, filter, update, &options.UpdateOptions{
		Upsert: &upsert,
	})
	if err != nil {
		log.Println("insert one error:", err)
	}
	log.Println("upsert result:", rs)

	find, err := conn.Find(ctx, filter)
	if err != nil {
		log.Println("insert one error:", err)
	}
	for find.Next(ctx) {
		log.Println(find.Current)
	}
}

// UpdateMany ...
func UpdateMany(ctx context.Context) {
	conn := mongoDBConnection(ctx)

	find := func(ctx context.Context) {
		rs, err := conn.Find(ctx, bson.D{{Key: "birthday", Value: "10/13/1978"}})
		if err != nil {
			log.Fatal(err)
		}
		for rs.Next(ctx) {
			log.Println(rs.Current)
		}
	}

	n := 5
	data := make([]any, 0, n)
	for i := 0; i < n; i++ {
		data = append(data, bson.D{{Key: "birthday", Value: "10/13/1978"}})
	}

	_, err := conn.InsertMany(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
	find(ctx)

	filter := bson.D{{Key: "birthday", Value: "10/13/1978"}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "gift", Value: "Happy BirthDay"}}}}
	_, err = conn.UpdateMany(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}
	find(ctx)
}

// Return updated documents
// Array operation
