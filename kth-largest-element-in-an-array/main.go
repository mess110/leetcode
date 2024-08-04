package main

import "fmt"

type MaxHeap struct {
	items []int
}

func (mh *MaxHeap) insert(e int) {
	mh.items = append(mh.items, e)
	mh.heapifyUp(len(mh.items) - 1)
}

// 1, 2, 3, 4, 5
func (mh *MaxHeap) heapifyUp(index int) {
	parentIndex := (index - 1) / 2

	if index > 0 && mh.items[index] > mh.items[parentIndex] {
		mh.items[index], mh.items[parentIndex] = mh.items[parentIndex], mh.items[index]
		mh.heapifyUp(parentIndex)
	}
}

func (mh *MaxHeap) remove() int {
	if len(mh.items) == 0 {
		return 0
	}

	root := mh.items[0]
	mh.items[0] = mh.items[len(mh.items)-1]
	mh.items = mh.items[:len(mh.items)-1]
	mh.heapifyDown(0)
	return root
}

func (mh *MaxHeap) heapifyDown(index int) {
	lastIndex := len(mh.items) - 1
	leftChildIndex := index*2 + 1
	rightChildIndex := index*2 + 2
	childToCompare := 0

	if leftChildIndex > lastIndex {
		return
	}

	if rightChildIndex > lastIndex {
		childToCompare = leftChildIndex
	} else if mh.items[leftChildIndex] > mh.items[rightChildIndex] {
		childToCompare = leftChildIndex
	} else {
		childToCompare = rightChildIndex
	}

	if mh.items[index] < mh.items[childToCompare] {
		mh.items[childToCompare], mh.items[index] = mh.items[index], mh.items[childToCompare]
		mh.heapifyDown(childToCompare)
	}
}

func findKthLargest(nums []int, k int) int {
	mh := MaxHeap{}
	for _, e := range nums {
		mh.insert(e)
	}

	var number int

	fmt.Println(mh.items)
	mh.remove()
	fmt.Println(mh.items)
	mh.remove()
	fmt.Println(mh.items)
	// for k > 0 {
	// 	number = mh.remove()
	// 	k--
	// }

	return number
}

func main() {
	nums := []int{3, 2, 1, 5, 6, 4, 12}
	k := 2
	// fmt.Println(findKthLargest(nums, k))

	// nums = []int{1}
	// k = 1
	// fmt.Println(findKthLargest(nums, k))

	nums = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k = 4
	fmt.Println("output", findKthLargest(nums, k))
}
