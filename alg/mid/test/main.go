package main

import (
	"Golearn/alg/mid"
	"fmt"
)

func main() {
	p := [][]int{{5, 3}}
	q := [][]int{{2, 3, 1}}
	fmt.Println(mid.CountPoints(p, q))
}
