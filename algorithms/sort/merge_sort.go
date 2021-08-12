package sortalgo

func mergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	var middle = len(arr) / 2
	return merge(mergeSort(arr[:middle]), mergeSort(arr[middle:]))
}

// solution 1
func merge(arr1 []int, arr2 []int) []int {
	var result = make([]int, len(arr1)+len(arr2))
	var i = 0
	var j = 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] <= arr2[j] {
			result[i+j] = arr1[i]
			i++
		} else {
			result[i+j] = arr2[j]
			j++
		}
	}

	// 两部分剩余未比较元素，直接追加到数组末尾（默认已排序）
	for i < len(arr1) {
		result[i+j] = arr1[i]
		i++
	}

	for j < len(arr2) {
		result[i+j] = arr2[j]
		j++
	}

	return result
}

func mergeSort2(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	var middle = len(arr) / 2
	return merge2(mergeSort2(arr[:middle]), mergeSort2(arr[middle:]))
}

// solution 2
func merge2(x []int, y []int) []int {
	// 初始化结果数组，存放序列化后的元素
	var r = make([]int, len(x)+len(y))
	for i, j := 0, 0; ; {
		// 数组 x 未消耗完毕 且 （数组 y 消耗完毕或者数组 x 值小于数组 y 时）
		if i < len(x) && (j == len(y) || x[i] < y[j]) {
			r[i+j] = x[i]
			i++
			// 数组 x 消耗完毕 或 （数组 y 未消耗完毕 且 数组 x 值大于数组 y 时）
		} else if j < len(y) {
			r[i+j] = y[j]
			j++
		} else {
			break
		}
	}

	return r
}
