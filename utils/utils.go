package utils

func IntsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true

}

func IntsCopy(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}

func IntSliceShorter(a, b []int) []int {
	if len(a) <= len(b) {
		return a
	}
	return b
}
