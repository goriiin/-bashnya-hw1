package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	_, err := fmt.Scan(&str)
	if err != nil {
		return
	}
	for strings.Index(str, "1") != -1 {
		str = str[:strings.Index(str, "1")] + "one" + str[strings.Index(str, "1")+1:]
	}

	fmt.Println(str)
}
