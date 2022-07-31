package main

import "fmt"

func main() {
	fmt.Println(test())
}

func test() int {
	ch := make(chan int, 3)
	go func() {
		ch <- 3
		fmt.Println(1)
	}()
	go func() {
		ch <- 4
		fmt.Println(2)
	}()
	go func() {
		ch <- 5
		fmt.Println(3)
	}()
	close(ch)
	return <-ch
}
