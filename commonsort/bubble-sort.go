package commonsort

/*
	冒泡排序的算法步骤

	1）比较相邻的元素。如果第一个比第二个大，就交换他们两个。
	2）对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
	3）针对所有的元素重复以上的步骤，除了最后一个。
	4）持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。

	不稳定排序，最好情况 O(n)，最差情况 O(n * n)
*/

func BubbleSort(n []int) []int {
	length := len(n)
	for i := 1; i < length; i++ {

		noNeedExchange := true

		for j := 0; j < length - i; j++ {
			if n[j] > n[j+1] {
				n[j], n[j+1] = n[j+1], n[j]
				noNeedExchange = false
			}
		}

		if noNeedExchange {
			break
		}
	}

	return n
}