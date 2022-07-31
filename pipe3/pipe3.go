package main

import "fmt"

func f1(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func f2(ch1 chan<- int, ch2 <-chan int) {
	for i := range ch2 {
		ch1 <- i
	}
	close(ch1)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go f1(ch2)
	go f2(ch1, ch2)
	for i := range ch1 {
		fmt.Println(i)
	}
}
