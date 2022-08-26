package main

import (
	"fmt"
	"sort"
)

type IntArray []int

func (i IntArray) Len() int {
	return len(i)
}

func (i IntArray) Less(ii, ij int) bool {
	return i[ii] < i[ij]
}

func (i IntArray) Swap(ii, ij int) {
	i[ii], i[ij] = i[ij], i[ii]
}

func main() {
	my := IntArray{5, 4, 6, 1, 2}
	sort.Sort(my)
	fmt.Println(my)
}
