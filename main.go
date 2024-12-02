package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

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

func matchExactWithWildCard(pattern, str string) bool {
	if len(pattern) != len(str) {
		return false
	}

	for i := 0; i < len(pattern); i++ {
		if pattern[i] != '.' && pattern[i] != str[i] {
			return false
		}
	}

	return true
}

func matchRegex(pattern, str string) bool {

	startMetachar := strings.HasPrefix(pattern, "^")
	if startMetachar {
		pattern = pattern[1:]
	}

	endMetachar := strings.HasSuffix(pattern, "$")
	if endMetachar {
		pattern = pattern[:len(pattern)-1]
	}

	if startMetachar && endMetachar {
		return matchExactWithWildCard(pattern, str)
	} else if startMetachar {
		str := str[:len(pattern)]
		return matchExactWithWildCard(pattern, str)
	} else if endMetachar {
		str := str[len(str)-len(pattern):]
		return matchExactWithWildCard(pattern, str)
	}

	for i := 0; i <= len(str)-len(pattern); i++ {
		if matchExactWithWildCard(pattern, str[i:i+len(pattern)]) {
			return true
		}
	}

	return false
}
