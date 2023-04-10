package main

import (
	"books-note/Mongodb-The-Definitive-Guide/chapter7"
	"context"
	"log"
	_ "net/http/pprof"
)

type Slice []bool

func (s Slice) Length() int {
	return len(s)
}
func (s Slice) Modify(i int, x bool) {
	s[i] = x // panic if s is nil
}
func (p *Slice) DoNothing() {}

func (p *Slice) Append(x bool) {
	if p == nil {
		p = new(Slice)
		v := []bool{x}
		*p = v
		return
	}
	*p = append(*p, x)
}

func main() {
	chapter7.Aggregate(context.Background())
}

type person struct{}

func (p person) Run() {
	log.Println("RUN")
}

type runner interface {
	Run()
}
