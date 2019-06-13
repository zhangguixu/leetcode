package commonsort

/*
	选择排序，算法步骤

	1）首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
	2）再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
	3）重复第二步，直到所有元素均排序完毕。

	稳定排序，时间复杂度：O(n * n)
*/

func SelectionSort(n []int) []int {
	l := len(n)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if n[i] > n[j] {
				n[i], n[j] = n[j], n[i]
			}
		}
	}
	return n
}