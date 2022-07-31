package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	done := make(chan struct{})
	go func() {
		time.Sleep(10 * time.Second)
		done <- struct{}{}
	}()

	for {
		select {
		case <-t.C:
			fmt.Println("1s")
		case <-done:
			fmt.Println("done")
		default:
			fmt.Println("default")
		}
	}
}
