package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go f1(ch1, ch2)
	go f2(ch2)
	for i := 0; i < 10; i++ {
		ch1 <- i
		time.Sleep(1 * time.Second)
	}
	close(ch1)
}

func f1(ch1, ch2 chan int) {
	for {
		num, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- num
	}

}

func f2(ch chan int) {
	for {
		num, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(num)
	}

}
