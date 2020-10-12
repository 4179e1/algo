package sort

/**
 * @brief partition for quick sort
 *
 * Select the last element as the Q value, elements between index p and i
 * are smaller than Q, elements between index i + 1 and r - 1 are greater than
 * or equal to Q. Swap Q with the elements in index (i + 1). return index of Q.
 *
 * @note Complexity O ( n )
 * @return Index of Q value.
 */

func Partition(a []int) int {
	var i, j, key int
	key = a[len(a)-1]

	i = 0 // end boundary for elements smaller than key
	for j = 0; j < len(a)-1; j++ {
		if a[j] < key {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[len(a)-1] = a[len(a)-1], a[i]

	return i
}

/**
 * @brief quick sort
 *
 * Partition array into 2 sub-sequences withe middle number Q index with q,
 * so that in each recursion, Q is placed in the right place.
 *
 * @note Complexity O ( n*lg(n) )
 */

func QuickSort(a []int) []int {
	var q int

	if len(a) > 1 {
		q = Partition(a)
		QuickSort(a[:q])
		QuickSort(a[q+1:])
	}

	return a
}

func Partition2(a []int, p, r int) int {
	key := a[r]

	i := 0
	for j := 0; j <= r; j++ {
		if a[j] < key {
			a[i], a[j] = a[j], a[i]
			i++
		}
	}
	a[i], a[r] = a[r], a[i]
	//fmt.Println(a)
	return i
}

// r -> len(a) - 1
func QuickSort2(a []int, p, r int) []int {
	if p < r {
		q := Partition2(a, p, r)
		QuickSort2(a, p, q-1)
		QuickSort2(a, q+1, r)

	}

	return a
}
