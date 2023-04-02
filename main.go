package main

import (
	"books-note/Mongodb-The-Definitive-Guide/chapter5"
	"context"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	chapter5.Indexes(ctx)
}
