package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(addBinary("11", "1")) // 100
	fmt.Println(addBinary("1010", "1011")) // 10101

	fmt.Println(addBinary_v1("11", "1")) // 100
	fmt.Println(addBinary_v1("1010", "1011")) // 10101
}

/*
	使用字符串的方式
*/
func addBinary(a string, b string) string {
	longStr, shortStr := a, b
	if len(a) < len(b) {
		longStr, shortStr = shortStr, longStr
	}
	i, j, idx := len(longStr) - 1, len(shortStr) - 1, 0
	sum := make([]byte, 0, len(longStr))
	var carry bool
	for j >= 0 {
		sum = append(sum, '0')
		switch longStr[i] - shortStr[j] {
		case 0:
			if carry {
				sum[idx] = '1'
			} else {
				sum[idx] = '0'
			}
			if longStr[i] == '1' {
				carry = true
			} else {
				carry = false
			}
		default:
			if carry {
				sum[idx] = '0'
				carry = true
			} else {
				sum[idx] = '1'
				carry = false
			}
		}
		idx++
		i--
		j--
	}
	for i >= 0 {
		sum = append(sum, '0')
		switch longStr[i] {
		case '0':
			if carry {
				sum[idx] = '1'
				carry = false
			}
		case '1':
			if !carry {
				sum[idx] = '1'
			}
		}
		i--
		idx++
	}
	if carry {
		sum = append(sum, '1')
	}
	reverse(sum)
	return string(sum)
}

func reverse(list []byte) {
	i, j := 0, len(list) - 1
	for i < j {
		list[i], list[j] = list[j], list[i]
		i++
		j--
	}
}

/*
	将字符串转为数字，然后数字进行运算之后，再转为二进制字符串
	
	在go语言中，可以使用strconv
	func ParseInt(s string, base int, bitSize int) (i int64, err error)
	func FormatInt(i uint64, base int) string

	字符串转为数字，会遇到溢出的问题，例如以下case就无法正确计算：
	fmt.Println(addBinary_v1("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101", "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011"))
*/
// func addBinary_v1(a string, b string) string {
// 	n1, _ := strconv.ParseInt(a, 64)
// 	n2, _ := strconv.ParseInt(b, 2, 64)
// 	return strconv.FormatInt(n1 + n2, 2)
// }