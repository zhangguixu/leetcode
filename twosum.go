package main

import (
	"fmt"
)


func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	res := make([]int, 0, 2)
	tmpMap := make(map[int][]int, len(nums))
    for index,n := range nums {
		list := tmpMap[n]
		if list == nil {
			list = make([]int, 0)
		}
		list = append(list, index)
		tmpMap[n] = list
	}
	for index,n := range nums {
		left := target - n
		if tmpMap[left] != nil {
			if left == n && len(tmpMap[left]) == 1 {
				continue
			}
			list := tmpMap[left]
			if left == n {
				res = list[0:2]
			} else {
				res = append(res, index)
				res = append(res, list[0])
			}
			break;
		}
	}
    return res
}

