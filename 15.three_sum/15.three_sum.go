package main

import "fmt"

func main() {
	arr1 := []int{-5, -3, -2, 1, 0, 3, 7, 8, 12, 15}
	data := threeSum(arr1)
	fmt.Println(data)
}

func threeSum(arr []int) [][]int {
	var result [][]int
	for index := 0; index < len(arr)-2; index++ {
		//skip duplicate data
		if index > 0 && arr[index] == arr[index-1] {
			continue
		}
		left := index + 1
		right := len(arr) - 1
		for left < right {
			sum := arr[index] + arr[left] + arr[right]
			if sum == 0 {
				result = append(result, []int{arr[index], arr[left], arr[right]})
				//skip duplicate data
				for left < right && arr[left] == arr[left+1] {
					left++
				}
				for left < right && arr[right] == arr[right-1] {
					right--
				}
				right--
				left++
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
