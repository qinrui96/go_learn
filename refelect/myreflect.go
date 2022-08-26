package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a = 3
	t := reflect.TypeOf(a)
	fmt.Println(t.Name(), t.Kind())
}
