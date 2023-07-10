package BSF

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
func isBadVersion(version int) bool {
	return true
}

func firstBadVersion(n int) int {
	//使用二分查找法
	l, r := 1, n
	for l <= r {
		mid := (r + l) / 2
		isBad := isBadVersion(mid)
		if !isBad {
			//向高版本测试
			l = mid + 1
		} else {
			//找到了或者向低版本测试
			if mid >= 1 && !isBadVersion(mid-1) {
				return mid
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
