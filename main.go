package main

import (
	"books-note/Mongodb-The-Definitive-Guide/chapter4"
	"context"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	// doc := chapter3.Document{
	// 	Title: "Hello MongoDB",
	// }
	// doc.InsertOne(ctx)
	// doc.InsertMany(ctx)
	// ordered := false
	// doc.InsertManyUnOrdered(ctx, ordered)
	// doc.DeleteOne(ctx, "642650fb7a709d119ef04cc3")
	// doc.Drop(ctx)

	// doc.UpdateOne(ctx, "64265af6db87c62981fda6b5")
	// doc.UpdateMany(ctx, "64265af6db87c62981fda6b5")
	// doc.SetOperator(ctx, "64265af6db87c62981fda6b5", "MongoDB")
	// doc.FindByID(ctx, "64265af6db87c62981fda6b5")
	// doc.SetEmbeddedDocument(ctx, "64265af6db87c62981fda6b5")
	// doc.FindByID(ctx, "64265af6db87c62981fda6b5")
	// doc.UnSetOperator(ctx, "64265af6db87c62981fda6b5", "name")
	// doc.FindByID(ctx, "64265af6db87c62981fda6b5")
	// doc.PushOperator(ctx)
	// doc.Upsert(ctx)
	// chapter3.UpdateMany(ctx)
	// chapter4.Find(ctx)
	// chapter4.Projection(ctx)
	// chapter4.QueryCondition(ctx)
	chapter4.OrQuery(ctx)
}
