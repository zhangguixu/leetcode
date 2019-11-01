package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse_v1(1234))
	fmt.Println(reverse_v1(-6789))
}

func reverse(x int) int {
	max := int(math.Pow(2, 31)) - 1
	min := int(-math.Pow(2, 31))
	hasMinus := false

	if x < 0 {
		hasMinus = true
		x = -x
	}

	num := make([]int, 0, 10)

	for x > 0 {
		num = append(num, x%10)
		x = x / 10
	}

	digits := len(num)
	x = 0

	for i := 0; i < digits; i++ {
		x += num[i] * int(math.Pow10(digits-i-1))
	}

	if hasMinus {
		x = -x
	}

	if x < min || x > max {
		x = 0
	}

	return x
}

/*
	从低位到高位依次获取一个数字的个位、十位、百位... 可以通过/和%操作，

	例如 123

	123 % 10 = 3

	123 / 10 = 12 	12 % 10 = 2

	12 / 10 = 1     1 % 10 = 1

	就可以依次获取[3, 2, 1]

	题目的要求是将数字反过来，在依次获取每个位数时，可以同时累加，

	digit                                 rev

	123 % 10 = 3                          3

	123 / 10 = 12 	12 % 10 = 2           3 * 10 + 2

	12 / 10 = 1     1 % 10 = 1            (3 * 10 + 2) * 10 + 1 = 321

	上述实现中，先将位数存储到了数组中，最后再进行累加，这里可以省去位数所占的空间

	这道题还有一个考察点：int的溢出问题，在discuss中，也有对应的讨论。

	https://leetcode.com/explore/featured/card/top-interview-questions-easy/127/strings/880/discuss/4060/My-accepted-15-lines-of-code-for-Java

	在代码中，使用的int，在go语言中，这个整数类型根据编译器和硬件平台的不同，有可能是32bit或者是64bit

	如果是int32，那么int类型表示的数字范围就是:
	
	int(math.Pow(2, 31)) - 1 = 2147483647  ~ int(-math.Pow(2, 31))  = -2147483648

	如果是int64，那么int类型表示的数字范围就是：

	-9223372036854775808 ~ 9223372036854775807

	这里如果要用discuss中的溢出判断的技巧，就必须指定int32类型
*/
func reverse_v1(x int) int {
	max := int(math.Pow(2, 31)) - 1
	min := int(-math.Pow(2, 31))

	var rev int

	for x != 0 {
		rev = rev * 10 + x % 10
		if rev < min || rev > max {
			rev = 0
			break
		}
		x = x / 10
	}

	return rev
}