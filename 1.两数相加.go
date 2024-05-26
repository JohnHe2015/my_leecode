package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 133
	result := addSum(nums, target)
	fmt.Println(result)
}

func addSum(arr []int, target int) []int {
	hashTable := make(map[int]int)
	for i, num := range arr {
		if j, ok := hashTable[target-num]; ok {
			return []int{j, i}
		}
		hashTable[num] = i
	}
	return nil
}
