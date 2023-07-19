package mid

import (
	"fmt"
	"math"
)

func CountPoints(points [][]int, queries [][]int) []int {
	//数学问题
	res := make([]int, len(queries))
	for j, q := range queries {
		for _, p := range points {
			if isContain(p, q) {
				res[j]++
			}
		}
	}
	return res
}

func isContain(p, quer []int) bool {
	//计算两点间距离
	x := (quer[0] - p[0]) * (quer[0] - p[0])
	y := (quer[1] - p[1]) * (quer[1] - p[1])
	fmt.Println("pow:", math.Pow(float64(x+y), 0.5))
	fmt.Println("r:", quer[2])
	return float64(quer[2]) >= math.Pow(float64(x+y), 0.5)
}

func countPointsV2(points [][]int, queries [][]int) []int {
	//数学问题
	res := make([]int, len(queries))
	for j, q := range queries {
		for _, p := range points {
			if isContainV2(p, q) {
				res[j]++
			}
		}
	}
	return res
}

func isContainV2(p, quer []int) bool {
	//计算两点间距离
	x := (quer[0] - p[0]) * (quer[0] - p[0])
	y := (quer[1] - p[1]) * (quer[1] - p[1])
	return quer[2]*quer[2] >= x+y
}
