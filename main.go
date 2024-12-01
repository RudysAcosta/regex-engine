package main

import (
	"fmt"
	"strings"
)

func main() {

	var input string

	fmt.Scan(&input)

	inputs := strings.Split(input, "|")

	pattern := inputs[0]
	expresion := inputs[1]

	if pattern == "" {
		fmt.Println(true)
	} else {
		if expresion == "" {
			fmt.Println(false)
		} else {
			fmt.Println(matchRegex(pattern, expresion))
		}
	}
}

func matchChar(pattern, char byte) bool {
	return pattern == char || pattern == '.'
}

func matchRegex(regex, str string) bool {
	i, j := 0, 0

	for i < len(str) && j < len(regex) {
		if matchChar(regex[j], str[i]) {
			i++
			j++
		} else if j > 0 {
			return false
		} else {
			i++
		}
	}

	return j == len(regex)
}
