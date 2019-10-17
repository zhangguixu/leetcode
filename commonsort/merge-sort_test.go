package commonsort

import "testing"

func Test_MergeSort(t *testing.T) {
	t.Log(MergeSort([]int{5, 3}))
	t.Log(MergeSort([]int{5, 3, 4, 1, 6, 8}))
	t.Log(MergeSort([]int{5, 3, 4, 1, 6, 8, 9}))
}
