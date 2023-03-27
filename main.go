package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{
		name: "tdn",
		age:  23,
	}
	p2 := person{
		name: "tdn",
		age:  23,
	}
	fmt.Println(p1 == p2)

}
