package main

import (
	"fmt"
)

func twoSum2(nums []int, target int) []int {
	for i, n := range nums {
		for j, m := range nums {
			if n+m == target && i != j {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

func twoSum(nums []int, target int) []int {
	adj := map[int]int{}
	fmt.Println("nums: ", nums)

	for i, n := range nums {
		_, exists := adj[target-n]
		if exists {
			return []int{i, adj[target-n]}
		}
		adj[n] = i
		fmt.Println("adj: ", adj)
	}
	// fmt.Println("adj: ", adj)

	return []int{}
}

func main() {
	nums := []int{}

	nums = []int{2, 7, 11, 15}
	fmt.Println(twoSum(nums, 9))

	nums = []int{3, 2, 4}
	fmt.Println(twoSum(nums, 6))

	nums = []int{3, 3}
	fmt.Println(twoSum(nums, 6))

	nums = []int{-3, 4, 3, 90}
	fmt.Println(twoSum(nums, 0))
}
