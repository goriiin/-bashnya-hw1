package main

import "fmt"

func main() {
	var (
		size int
		num  int
	)
	_, err := fmt.Scan(&size)
	if err != nil {
		return
	}

	arr := make([]int, size)
	for i := 0; i < size-1; i++ {
		_, err = fmt.Scan(&num)
		if err != nil {
			return
		}
		arr[i+1] = num
	}

	_, err = fmt.Scan(&num)
	if err != nil {
		return
	}
	arr[0] = num
	fmt.Println(arr)
}
