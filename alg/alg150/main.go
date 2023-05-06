package main

import (
	"Golearn/alg/alg150/alg/SlidingWindow"
	"fmt"
)

func main() {

	//fmt.Println(alg.AergeAlternately("abcABC", "def"))
	//fmt.Println(alg.CanPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2))
	//fmt.Println(alg.ReverseVowels("leetcode"))
	//fmt.Println(ArrayAndString.ReverseWords("    the    sky is    blue    "))
	//fmt.Println(ArrayAndString.ProductExceptSelf2([]int{-1, 1, 0, -3, 3}))
	//fmt.Println(ArrayAndString.Compress([]byte{'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'c', 'b'}))
	//DoublePointer.MoveZeroes1([]int{0, 1, 0, 3, 12})
	//fmt.Println(DoublePointer.IsSubsequence("abc", "ahbgdc"))
	//fmt.Println(DoublePointer.MaxArea1([]int{1, 2, 4, 3}))

	//arr := []int{1, 2, 4, 5}
	//t := time.Now()
	//fmt.Println(DoublePointer.MaxOperations(arr, 5))
	//fmt.Println(time.Now().Sub(t))
	//t = time.Now()
	//fmt.Println(DoublePointer.MaxOperations1(arr, 5))
	//fmt.Println(time.Now().Sub(t))

	fmt.Println(SlidingWindow.MaxVowels("leetcode", 3))

}
