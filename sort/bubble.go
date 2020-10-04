package sort

/**
 * @brief Bubble Sort for array.
 *
 * Repeatedly swapping adjacent elements that are out of order.
 *
 * @note Complexity: O(n^2).
 * @param array The array need to be sort.
 * @param len Length of array.
 * @return The sorted array.
 */

func BubbleSort(a []int) []int {
	var i, j int

	for i = 0; i < len(a)-1; i++ {
		for j = len(a) - 1; j > i; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}

	return a
}
