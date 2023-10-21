package main

import "fmt"

func check(a, b, c int) bool {
	return a+b > c
}

func triangleExists(a, b, c int) string {
	if a < 0 || b < 0 || c < 0 {
		return "NO"
	}

	if check(a, b, c) && check(a, c, b) && check(b, c, a) {
		return "YES"
	}

	return "NO"
}

func main() {
	fmt.Println(triangleExists(3, 4, 5))
}
