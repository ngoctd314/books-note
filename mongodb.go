package main

import (
	"books-note/Mongodb-The-Definitive-Guide/chapter3"
	"context"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()

	doc := chapter3.Document{
		Title: "Hello MongoDB",
	}
	// doc.InsertOne(ctx)
	doc.InsertMany(ctx)
}
