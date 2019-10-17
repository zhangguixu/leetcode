package commonsort

import "testing"

func Test_QuickSort(t *testing.T) {
	t.Log(QuickSort([]int{5, 3}))
	t.Log(QuickSort([]int{5, 3, 4, 1, 6, 8}))
	t.Log(QuickSort([]int{5, 3, 4, 1, 6, 8, 9}))
}
