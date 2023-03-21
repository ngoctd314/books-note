package main

import (
	"fmt"
	"reflect"
)

type T struct {
	X    int  `max:"99" min:"0" default:"0"`
	Y, Z bool `optional:"yes"`
}

func main() {
	t := reflect.TypeOf(T{})
	x := t.Field(0).Tag
	y := t.Field(1).Tag
	z := t.Field(2).Tag

	fmt.Println(x, y, z)
	fmt.Println(reflect.TypeOf(x))
	v, present := x.Lookup("max")
	fmt.Println(len(v), present, v)
	fmt.Println(x.Get("max"))
	fmt.Println(x.Lookup("optional"))
	fmt.Println(y.Lookup("optional"))
	fmt.Println(z.Lookup("optional"))
}
