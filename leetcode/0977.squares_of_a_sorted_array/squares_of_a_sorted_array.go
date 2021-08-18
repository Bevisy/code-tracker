package leetcode

import "sort"

// 双指针
func sortedSquares(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	i, j := 0, n-1
	for pos := n - 1; pos >= 0; pos-- {
		if v, w := nums[i]*nums[i], nums[j]*nums[j]; v > w {
			ans[pos] = v
			i++
		} else {
			ans[pos] = w
			j--
		}
	}
	return ans
}

// 双指针
func sortedSquares1(nums []int) []int {
	n := len(nums)
	lastNegIndex := -1
	for i := 0; i < n && nums[i] < 0; i++ {
		lastNegIndex = i
	}

	ans := make([]int, 0, n)
	for i, j := lastNegIndex, lastNegIndex+1; i >= 0 || j < n; {
		if i < 0 {
			ans = append(ans, nums[j]*nums[j])
			j++
		} else if j == n {
			ans = append(ans, nums[i]*nums[i])
			i--
		} else if nums[i]*nums[i] < nums[j]*nums[j] {
			ans = append(ans, nums[i]*nums[i])
			i--
		} else {
			ans = append(ans, nums[j]*nums[j])
			j++
		}
	}

	return ans
}

// 时间复杂度： O(n logn)
func sortedSquares2(nums []int) []int {
	for i := range nums {
		nums[i] = nums[i] * nums[i]
	}

	sort.Ints(nums)
	return nums
}

// sort.Ints() 内存占用和速度远远优于 quickSort() 实现
func quickSort(array []int) []int {
	if len(array) < 1 {
		return array
	}

	var smaller, greater []int
	pivot := 0
	for i := 1; i < len(array); i++ {
		if array[i] < array[pivot] {
			smaller = append(smaller, array[i])
		} else {
			greater = append(greater, array[i])
		}
	}

	r := append(quickSort(smaller), array[pivot])
	r = append(r, quickSort(greater)...)
	return r
}
