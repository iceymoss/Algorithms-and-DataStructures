package DoublePointer

//暴力枚举：找出最大值
//时间：O(n^2); 空间：O(1)

func MaxArea(height []int) int {
	maxA := 0
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			maxA = max(maxA, (j-i)*min(height[i], height[j]))
		}
	}
	return maxA
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

//双指针法: 左右指针，每次找出一个容量值，然后选择移动边界小的值(移动短板),
//如果不移动短板：加入左指针是短板，我们不移动他的话，我们移动右指针，那么短板永远存在，
//并且我们在右指针是向左指针移动的，得到的结果只有两种可能小于或者等于移动指针前的数

func MaxArea1(height []int) int {
	if len(height) <= 1 {
		return 0
	}
	maxA := 0
	left, right := 0, len(height)-1
	for left < right {
		maxA = max(maxA, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxA
}
