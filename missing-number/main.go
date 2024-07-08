package main

import (
	"fmt"
	"sort"
)

func missingNumber(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	j := 0
	for i := range len(nums) {
		j = i
		if nums[i] != i {
			return i
		}
	}
	return j + 1
}

/* interesting alternative solution:
total_sum = sum of n numbers = n(n+1) / 2
nums_sum = sum(nums)
return total_sum - nums_sum
*/

func main() {
	nums := []int{}

	nums = []int{3, 0, 1}
	fmt.Println(missingNumber(nums) == 2)

	nums = []int{0, 1}
	fmt.Println(missingNumber(nums) == 2)

	nums = []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(nums) == 8)

	nums = []int{1}
	fmt.Println(missingNumber(nums) == 0)

	nums = []int{0}
	fmt.Println(missingNumber(nums) == 1)

	nums = []int{1, 2}
	fmt.Println(missingNumber(nums) == 0)
}
