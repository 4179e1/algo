package sort

func InsertSort(a []int) []int {
	/**
	 * j for the select card in array
	 * i for the boundary of the sorted group in the beginging
	 * i + 1 for the index for the selected card after iteration.
	 * key for the value of the selected card.
	**/

	var i, j, key int

	for j = 1; j < len(a); j++ {
		key = a[j]
		i = j - 1
		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i--
		}
		a[i+1] = key
	}

	return a
}
