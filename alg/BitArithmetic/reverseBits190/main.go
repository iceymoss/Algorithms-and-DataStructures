package main

import "fmt"

// 本题使用位运算，需要搞清楚什么是左右移位，取最后一位值等
func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		//判断最后一位是否为1
		isZore := num & 1
		//num右移一位
		num >>= 1
		//res左移一位,这里一共会左移 32 次
		res <<= 1
		if isZore == 1 {
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(reverseBits(1111010))
}
