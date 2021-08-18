package leetcode

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	// 处理数组数据极端情况，均大于或者均小于 target
	if nums[left] > target {
		return 0
	} else if nums[right] < target {
		return right + 1
	}

	// target 位于数组内部， 使用二分法寻找确切位置
	for left < right {
		pivot := left + (right-left)/2
		if nums[pivot] == target {
			return pivot
		} else if nums[pivot] < target {
			left = pivot + 1
		} else {
			right = pivot
		}
	}

	return left
}
