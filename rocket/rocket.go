package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		buf := make([]byte, 1)
		os.Stdin.Read(buf)
		abort <- struct{}{}
	}()

	tch := time.Tick(1 * time.Second)

	for i := 0; i < 10; i++ {
		select {
		case <-tch:
		case <-abort:
			fmt.Println("abort")
			return
		}
	}
	fmt.Println("lunch")
}
