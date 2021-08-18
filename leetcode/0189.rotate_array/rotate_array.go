package leetcode

// 环状替换
func rotate(nums []int, k int) {
	
}

// 利用额外的数组，完成翻转
// 时间复杂度 O(n)
// 空间复杂度 O(n)
func rotate2(nums []int, k int) {
	// 超过数组长度整数倍的位移视为无效操作
	n := len(nums)
	// k = k % n

	newNums := make([]int, n)
	for i := range nums {
		newNums[(k+i)%n] = nums[i]
	}
	copy(nums, newNums)
}
