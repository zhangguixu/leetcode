package commonsort

import "testing"

func Test_SelectionSort(t *testing.T) {
	t.Log(SelectionSort([]int{5, 3, 4, 1, 6, 8}))
}
