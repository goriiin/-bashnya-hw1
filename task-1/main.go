package main

import (
	"fmt"
	"math"
)

type MyError struct {
	message string
}

func (e *MyError) Error() string {
	return e.message
}

func findHypotenuse(a, b int) (float64, error) {
	if a < 0 || b < 0 {
		return -1, &MyError{"negative nums"}
	}

	return math.Sqrt(float64(a*a + b*b)), nil
}

func main() {
	fmt.Println(findHypotenuse(3, 4))
}
