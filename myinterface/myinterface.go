package main

import "fmt"

type A interface {
	haha()
}

type Ta struct {
	a int
}

func (t Ta) haha() {
	fmt.Println(t.a)
}

func test(a A) {
	a.haha()
}

func main() {
	mm := &Ta{}
	test(mm)
}
