/*
真题描述：给定一个非空字符串 s，最多删除一个字符。判断是否能成为回文字符串。

示例 1: 输入: "aba"
输出: True
示例 2:
输入: "abca"
输出: True
解释: 你可以删除c字符。
注意: 字符串只包含从 a-z 的小写字母。字符串的最大长度是50000。
*/

// 快慢指针
package main

import "fmt"

func main() {
	fmt.Println(validPalindrome2("aba"))
	fmt.Println(validPalindrome2("abca"))
	fmt.Println(validPalindrome2("hello"))
}

func validPalindrome2(str string) bool {
	isPal := func(left, right int) bool {
		for left < right {
			if str[left] != str[right] {
				return false
			}
			left++
			right--
		}
		return true
	}

	left, right := 0, len(str)-1
	for left < right {
		if str[left] != str[right] {
			//如果遇到不同的情况，则有一次机会删除一个字符，然后在判断是否是回文
			return isPal(left+1, right) || isPal(left, right-1)
		}
		left++
		right--
	}

	return true
}
