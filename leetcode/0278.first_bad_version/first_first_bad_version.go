package leetcode

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

var arr []bool

func isBadVersion(version int) bool {
	return !arr[version-1]
}

// 暴力遍历：返回第一个 false
func firstBadVersion1(n int) int {
	for i := 1; i <= n; i++ {
		if isBadVersion(i) {
			return i
		}
	}
	return 0
}

// 二分法
// 通过二分法寻找 left == true && right == false 条件
func firstBadVersion(n int) int {
	l := 1
	r := n

	// 循环条件 l < r，最终 l == r 代表同一个点，结束循环
	for l < r {
		m := (l + r) / 2
		if isBadVersion(m) {
			r = m // 保证右边为false
		} else {
			l = m + 1 // 左边为 true，逼近右边
		}

	}

	return l
}
