package leetcode

// 双指针，原地修改
// 左指针指向已处理好的序列尾部，右指针指向待处理序列头部
func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		// 右指针每次循环必须移动一位，因为每次循环都有一个值被确认位置
		right++
	}
}

// 借助额外的数组空间
// 时间复杂度 O(n)
func moveZeroes2(nums []int) {
	newNums := make([]int, len(nums))
	i := 0
	for j := range nums {
		if nums[j] != 0 {
			newNums[i] = nums[j]
			i++
		}
	}

	copy(nums, newNums)
}
