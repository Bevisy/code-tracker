package sortalgo

// BubbleSort 冒泡排序. data必须实现sort包中的Interface接口
//func bubbleSort(data sort.Interface) {
//	n := data.Len()
//	for i := 0; i < n-1; i++ {
//		isChanged := false
//		for j := 0; j < n-1-i; j++ {
//			if data.Less(j, j+1) {
//				data.Swap(j, j+1)
//				isChanged = true
//			}
//		}
//		if !isChanged {
//			break
//		}
//	}
//}

// 方法二：从数组头部开始遍历，每次内循环，最大的数下沉到底
// 冒泡排序时间复杂度：平均 O(n^2); 最佳 O(n); 最差 O(n^2)
func bubbleSort2(array []int) []int {
	// 用来计数已排好序的元素个数，大数下沉，从尾部开始排序
	for i := 0; i < len(array)-1; i++ {
		// 标志位：确认本轮是否存在数据变动，不存在，则数组排序完成，提前结束
		isChanged := false
		// 判断条件 j > len(array)-1-i，从数组头部开始，每次遍历确认一个最大的数
		for j := 0; j < len(array)-1-i; j++ {
			if array[j+1] < array[j] {
				array[j], array[j+1] = array[j+1], array[j]
				isChanged = true
			}
		}

		if !isChanged {
			return array
		}
	}

	return array
}

// 方法三：从数组尾部开始遍历，每次内循环，最小的数上浮到顶
func bubbleSort3(array []int) []int {
	// 用来计数已排好序的元素个数，小数上浮，从头部开始排序
	for i := 0; i <= len(array)-1; i++ {
		var isChanged = false
		// 判断条件 j > i，从数组尾部开始，每次遍历确认一个最小的数
		for j := len(array) - 1; j > i; j-- {
			if array[j] < array[j-1] {
				array[j], array[j-1] = array[j-1], array[j]
				isChanged = true
			}
		}

		if !isChanged {
			return array
		}
	}
	return array
}
