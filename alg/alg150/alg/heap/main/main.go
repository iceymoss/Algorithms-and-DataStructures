package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(h.FindKthLargest([]int{2, 35, 10, 0, 110, 200}, 1))
	fmt.Println(sort.SearchInts([]int{2, 35, 100, 200, 210, 300}, 36))
}
