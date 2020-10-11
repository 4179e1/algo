package utils

import "bytes"

import "fmt"

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

func IntMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func IntsCopy(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}

func IntSliceShorter(a, b []int) []int {
	if len(a) < len(b) {
		return a
	}
	return b
}

func NewIntMatrix(m, n int) [][]int {
	a := make([][]int, m)
	for i := range a {
		a[i] = make([]int, n)
	}
	return a
}

type IntMatrix [][]int

func (m IntMatrix) String() string {
	var buffer bytes.Buffer

	//buffer.WriteRune('\t')
	for i := range m[0] {
		buffer.WriteString(fmt.Sprintf("\tCol %d", i))
	}
	buffer.WriteRune('\n')

	for i := range m {
		buffer.WriteString(fmt.Sprintf("Row %d:\t", i))
		for j := range m[i] {
			buffer.WriteString(fmt.Sprintf("%d\t", m[i][j]))
		}
		buffer.WriteRune('\n')
	}

	return buffer.String()
}
