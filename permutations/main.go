package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(permute(nums))
}

func permute(nums []int) [][]int {
	result := [][]int{}

	var backtrack func([]int, []int)
	backtrack = func(nums []int, path []int) {
		if len(nums) == 0 {
			result = append(result, append([]int{}, path...))
		}
		for i := range nums {
			newNums := append([]int{}, nums[:i]...)
			newNums = append(newNums, nums[i+1:]...)
			newPath := append(path, nums[i])
			backtrack(newNums, newPath)
		}
	}

	backtrack(nums, []int{})

	return result
}
