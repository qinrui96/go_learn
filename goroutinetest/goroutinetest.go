package main

import (
	"fmt"
	"time"
)

func f1() {
	go func() {
		for {
			fmt.Println("1")
			time.Sleep(1 * time.Second)
		}
	}()
	fmt.Println("f1 exit")
}

func main() {
	go f1()
	for {

	}
}
