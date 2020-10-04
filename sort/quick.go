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
/*
int partition (int *array, int p, int r)
{
	int i, j, x;
	x = array[r];
	i = p - 1;

	for (j = p; j &lt; r; j++)
	{
		if (array[j] &lt; x)
		{
			i++;
			swap (&amp;array[j], &amp;array[i]);
		}
	}

	swap (&amp;array[i + 1], &amp;array[r]);
	return i + 1;
}
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
