package main

import (
	"fmt"
	"strings"
)

func main() {
	m := make(map[string]bool)

	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		return
	}

	arr := strings.FieldsFunc(str, func(r rune) bool {
		return r == ' '
	})

	for i := 0; i < len(arr); i++ {
		m[arr[i]] = true
	}

	fmt.Println(len(m))
}
