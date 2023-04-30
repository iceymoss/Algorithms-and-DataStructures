package alg

//题目要求是否能在不打破规则的情况下插入n朵花，与直接计算不同，采用“跳格子”的解法只需遍历不到一遍数组，处理以下两种不同的情况即可：
//【1】当遍历到index遇到1时，说明这个位置有花，那必然从index+2的位置才有可能种花，因此当碰到1时直接跳过下一格。
//【2】当遍历到index遇到0时，由于每次碰到1都是跳两格，因此前一格必定是0，此时只需要判断下一格是不是1即可得出index这一格能不能种花，
//如果能种则令n减一，然后这个位置就按照遇到1时处理，即跳两格；如果index的后一格是1，说明这个位置不能种花且之后两格也不可能种花（参照【1】），直接跳过3格。
//当n减为0时，说明可以种入n朵花，则可以直接退出遍历返回true；如果遍历结束n没有减到0，说明最多种入的花的数量小于n，则返回false。

func CanPlaceFlowers(flowerbed []int, n int) bool {
	//遍历数组

	for i := 0; i < len(flowerbed) && n > 0; {
		//根据种花规则，当i位置有花时，一定要挑两格
		if flowerbed[i] == 1 {
			i += 2

			//当i位置没有种花时，根据i+=2可以说明i-1位置一定没有种花，只需要判断i+1位置是否有花
		} else if i == len(flowerbed)-1 || flowerbed[i+1] == 0 {
			n--
			i += 2
			//当i+1有花，我们需要跳过i+1,然后再跳两格，即i+=3
		} else {
			i += 3
		}
	}
	return n <= 0
}
