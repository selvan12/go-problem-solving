package main

/*
Given an m x n matrix mat, return an array of all the elements of the array in a diagonal order.

Example 1:
Input: mat = [[1,2,3],[4,5,6],[7,8,9]]
Output: [1,2,4,7,5,3,6,8,9]

Example 2:
Input: mat = [[1,2],[3,4]]
Output: [1,2,3,4]
*/
func findDiagonalOrder(mat [][]int) []int {
	diagonal := []int{}

	//height, width := len(mat), len(mat[0])
	//fmt.Printf("Height: %d, Width: %d\n", height, width)

	diagonals := make([][]int, len(mat)+len(mat[0]))
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			diagonals[i+j] = append(diagonals[i+j], mat[i][j])
		}
	}
	//fmt.Println("Diagonals: ", diagonals)

	for ind, val := range diagonals {
		if ind%2 == 0 {
			reverse_array(val)
		}
		diagonal = append(diagonal, val...)
	}

	return diagonal
}

func reverse_array(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
