package main

import "fmt"

//位运算：1、n & (-n) 可以得到 n 二进制式最后一位 1
//		2、n & (n-1) 打掉n最后一个1

func hammingWeight(num uint32) int {
	var res int
	for num != 0 {
		//num &= num - 1打掉最后一个1
		num &= num - 1
		res++
	}
	return res
}

// num & 1与运算，判断最后一位是否为1，num >> 1向右移
func hammingWeight1(num uint32) int {
	var res int
	for i := 0; i < 32; i++ {
		res += int(num & 1)
		num = num >> 1
	}
	return res
}

func denTobin(n int) []int {
	stack := make([]int, 0)
	for n > 0 {
		stack = append(stack, n%2)
		n = n / 2
	}
	return stack
}

func main() {
	fmt.Println(denTobin(5))
}
