package main

import "fmt"

func setZeroes(matrix [][]int) {
	// fmt.Println("maxtrix:", matrix)

	zeroes := [][]int{}
	for i, row := range matrix {
		for j, p := range row {
			if p == 0 {
				zeroes = append(zeroes, []int{i, j})
			}
		}
	}

	// fmt.Println("zeroes: ", zeroes)

	for _, value := range zeroes {
		matrix[value[0]] = make([]int, len(matrix[0]))
		for i, _ := range matrix {
			matrix[i][value[1]] = 0
		}
	}
}

func main() {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	fmt.Println(matrix)

	matrix = [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	fmt.Println(matrix)
}
