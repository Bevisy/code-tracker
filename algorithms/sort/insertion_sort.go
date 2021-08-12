package sortalgo

//插入排序
//算法时间复杂度O(n^2)
// func insertionSort(array []int) []int {
// 	for currentIndex := 1; currentIndex < len(array); currentIndex++ {
// 		temp := array[currentIndex]
// 		iterator := currentIndex
// 		//已排序的数组，从后往前检索，直到找到需插入数据的合适位置
// 		for ; iterator > 0 && array[iterator-1] >= temp; iterator-- {
// 			array[iterator] = array[iterator-1]
// 		}
// 		array[iterator] = temp
// 	}
// 	return array
// }

// 下标 i 为需要排序的元素
func insertionSort2(array []int) []int {
	// 从数组第二位开始，前面部分默认已排序
	for i := 1; i < len(array); i++ {
		// 从已排序部分尾部开始比较
		for j := i - 1; j >= 0; j-- {
			if array[i] < array[j] {
				array[i], array[j] = array[j], array[i]
				i = j // i,j 交换，目标 array[i] 代表的值下标随之变化；此处的 i 为局部变量
			}
		}
	}
	return array
}
