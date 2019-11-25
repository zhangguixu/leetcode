package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("42"))  // 42
	fmt.Println(myAtoi(" 	-42")) // -42
	fmt.Println(myAtoi("4193 with words")) // 4193
	fmt.Println(myAtoi("words and 987")) // 0
	fmt.Println(myAtoi("-91283472332")) // -2147483648
	fmt.Println(myAtoi("+122")) // 122
	fmt.Println(myAtoi("     +004500")) // 4500
	fmt.Println(myAtoi("   -04f")) // -4
	fmt.Println(myAtoi("   +0 123")) // 0
	fmt.Println(myAtoi("-5-")) // -5
	fmt.Println(myAtoi("-5+")) // -5
	fmt.Println(myAtoi("9223372036854775808")) // 2147483647
	fmt.Println(myAtoi("+-2")) // 0
	fmt.Println(myAtoi("2147483648")) // 2147483647
	fmt.Println(myAtoi("-21474836489")) // -2147483648
	fmt.Println(myAtoi("-6147483648")) // -2147483648
	fmt.Println(myAtoi("2147483646")) // 2147483646
}


/*
	ascii码表 http://ascii.911cha.com/

	本人理解，此题目的考察点

	1. 了解ascii码
	2. 对于异常逻辑的考虑是否周全
		1) 各个字符串的情况
		2) 溢出处理
	3. 实现是否简洁

	一些基础的知识

	整数在计算机中的表示

	1）原码：一个整数，按照绝对值大小转换成的二进制数，称为原码。
	2）反码：将二进制数按位取反，所得的新二进制数称为原二进制数的反码。
	3）补码：反码加1称为补码。

	正整数在计算机中是以原码的形式表示，负整数是以补码的形式表示

	go语言中，由于编译器不同或者硬件平台不同，int类型可能是int32，也可以是int64

	int32的取值范围是：math.Pow(2, 31) - 1 ~ -math.Pow(2, 31)

	判断整数溢出，有以下的思路：
		
		1）指定num为double，可以存储比int32更大的数，leetcode的测试用例最大没有超过double取值范围
		2) 指定num为int32，见代码的实现:
			- 乘法运算溢出的判断
				num * 10 / 10 != num 证明溢出
			- 加法运算溢出的判断，有z = x + y
				若x > 0, y > 0 且 z < 0，证明溢出
				若x < 0, y < 0 且 z > 0, 证明溢出
*/
const (
	Space = 32 // 空格
	Tab   = 9  // tab键
	Minus = 45 // -号
	Plus  = 43 // +号
)

func myAtoi(str string) int {
	var num int32
	var toIntStart bool
	symbol := int32(1)
	min := -int(math.Pow(2, 31))
	max := int(math.Pow(2, 31)) - 1
	extreme := max
	for i := 0; i < len(str); i++ {
		b := str[i]
		switch {
		case b == Minus:
			if toIntStart {
				return int(num)
			}
			toIntStart = true
			symbol = -1
			extreme = min
		case b == Space || b == Tab || b == Plus:
			if toIntStart {
				return int(num)
			}
			if b == Plus {
				toIntStart = true
			}
		case b >= 48 && b <= 57:
			toIntStart = true
			digit := int32(b - 48)
			// 检测数字是否溢出
			if num * 10 / 10 != num || (symbol == 1 && num * 10 + symbol * digit < 0) || (symbol == -1 && num * 10 + symbol * digit > 0) {
				return extreme
			}
			num = num * 10 + symbol * digit
		default:
			return int(num)
		}
	}
	return int(num)
}