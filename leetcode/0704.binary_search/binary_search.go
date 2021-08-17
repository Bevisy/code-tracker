package leetcode

// 时间复杂度：O(logN)
// 空间复杂度：O(1)
func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	// 注意：
	// 	golang 中 while 循环方式实现，省略初始条件和循环递增操作，并且不加 “;”
	for left <= right {
		middle := (left + right) / 2 // middle 的负值在条件判断后
		if nums[middle] == target {
			return middle
		} else if nums[middle] < target {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}
