package main

import "fmt"

// isPowerOfTwo 方法一能被二整除
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	//被2整除
	for n > 1 {
		if n%2 != 0 {
			return false
		}
		n = n / 2
	}
	return true
}

//方法二：使用库函数

// 方法三：位运算
func isPowerOfTwo2(n int) bool {
	//对于是2的n次方的数的二进制一定会满足：
	//二进制数里面只有一个1，一定不以1结尾
	return n > 0 && n&(n-1) == 0
}

func main() {
	fmt.Println(isPowerOfTwo2(1))

}
