package main

import "fmt"

func main() {
	var A interface{}
	A = 3
	switch A := A.(type) {
	case int:
		fmt.Println(A)
	}
}
