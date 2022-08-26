package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var val interface{}
	var num int
	val = num
	fmt.Println(unsafe.Sizeof(val))

	var bo bool
	val = bo
	fmt.Println(unsafe.Sizeof(val))
}
