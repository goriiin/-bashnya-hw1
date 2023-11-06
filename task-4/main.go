package main

import "fmt"

func createMatrix(size int) [][]int {
	arr := make([]int, size*size)

	matrix := make([][]int, size)
	for i := range matrix {
		matrix[i], arr = arr[:size], arr[size:]
	}
	return matrix
}

func fillMatrix(arr *[][]int) {
	for i := range *arr {
		var num int
		for j := range *arr {
			_, err := fmt.Scan(&num)
			if err != nil {
				return
			}
			(*arr)[i][j] = num
		}
	}
}

func main() {
	var size int
	_, err := fmt.Scan(&size)
	if err != nil {
		return
	}

	matrix := createMatrix(size)
	fillMatrix(&matrix)

	flag := true

	for i := range matrix {
		for j := range matrix {
			if matrix[i][j] != matrix[j][i] {
				flag = false
			}
		}
	}

	if flag {
		fmt.Println(`YES`)
	} else {
		fmt.Println("NO")
	}

}
