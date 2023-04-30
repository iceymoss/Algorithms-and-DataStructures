package main

import "fmt"

func numTrees(n int) int {
	//dp[i]表示1->i的二分搜索树的总数
	dp := make([]int, n+1)
	dp[0] = 1
	//以i为根节点
	for i := 1; i <= n; i++ {
		//计算以当前i为根节点构成的bfs
		for j := 1; j <= i; j++ {
			//在上⾯的分析中，其实已经看出其递推关系， dp[i] += dp[以j为头结点左⼦树节点数量] *
			//dp[以j为头结点右⼦树节点数量] j相当于是头结点的元素，从1遍历到i为⽌。
			//所以递推公式：dp[i] += dp[j - 1] * dp[i - j]; ，j-1 为j为头结点左⼦树节点数量，i-j 为以j为
			//头结点右⼦树节点数量
			dp[i] += dp[j] * dp[i-j]
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(numTrees(3))
}
