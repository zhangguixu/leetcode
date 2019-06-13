package commonsort

/*
	归并排序，算法步骤

	1）申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列；
	2）设定两个指针，最初位置分别为两个已经排序序列的起始位置；
	3）比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置；
	4）重复步骤 3 直到某一指针达到序列尾；
	5）将另一序列剩下的所有元素直接复制到合并序列尾

	可以用于外存排序，稳定排序，
	时间复杂度度为 O(nlogN)
	空间复杂度为O(n)
*/

func MergeSort(n []int) []int {
	if len(n) <= 1 {
		return n
	}
	middle := len(n) / 2
	return merge(MergeSort(n[:middle]), MergeSort(n[middle:]))
}

func merge(left []int, right []int) []int {
	list := make([]int, 0, len(left))
	var k, j int
	for k < len(left) && j < len(right) {
		if left[k] < right[j] {
			list = append(list, left[k])
			k++
		} else {
			list = append(list, right[j])
			j++
		}
	}
	for k < len(left) {
		list = append(list, left[k])
		k++
	}
	for j < len(right) {
		list = append(list, right[j])
		j++
	}
	return list
}
