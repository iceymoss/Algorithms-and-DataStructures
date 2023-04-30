package main

import "fmt"

func countBits(n int) []int {
	res := make([]int, n+1) //返回结果需要n+1的空间
	for i := 0; i <= n; i++ {
		//统计每一个i的值的二进制1的个数
		res[i] = hammingWeight(uint32(i))
	}
	return res
}

// hammingWeight 计算num的二进制数中的1的个数
func hammingWeight(num uint32) int {
	var res int
	for num != 0 {
		num &= num - 1
		res++
	}
	return res
}

func main() {
	fmt.Println(countBits(5))
}
