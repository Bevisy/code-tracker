package sortalgo

// out-place
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

// in-place
func quickSort2(array []int) []int {
	// TODO: in-palce 实现 https://zh.wikipedia.org/wiki/%E5%BF%AB%E9%80%9F%E6%8E%92%E5%BA%8F#%E5%8E%9F%E5%9C%B0%EF%BC%88in-place%EF%BC%89%E5%88%86%E5%89%B2%E7%9A%84%E7%89%88%E6%9C%AC
	return []int{}
}
