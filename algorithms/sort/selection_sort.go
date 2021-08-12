package sortalgo

//选择排序
// 选择排序时间复杂度：平均 O(n^2); 最佳 O(n^2); 最差 O(n^2)
func selectionSort(array []int) []int {
	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	return array
}
