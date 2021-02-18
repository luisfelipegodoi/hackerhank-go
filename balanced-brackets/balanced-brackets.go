package main

import (
	"fmt"
	"strings"
)

func isBalanced(s string) string {

	if s == "" {
		return "NO"
	}

	isPar := len(s) % 2
	if isPar != 0 {
		return "NO"
	}

	arrBrackets := []string{}
	arrS := strings.Split(s, "")
	for _, v := range arrS {

		if v == "{" || v == "[" || v == "(" {
			arrBrackets = append(arrBrackets, v)
			continue
		}

		switch v {
		case ")":
			if arrBrackets[2] == "{" || arrBrackets[2] == "[" {
				return "NO"
			}
			break
		case "}":
			if arrBrackets[0] == "(" || arrBrackets[0] == "[" {
				return "NO"
			}
			break
		case "]":
			if arrBrackets[1] == "(" || arrBrackets[1] == "{" {
				return "NO"
			}
			break
		}

	}

	return "YES"
}

func main() {

	var inputBrackets string = "{[(]]}"

	result := isBalanced(inputBrackets)

	fmt.Println(result)
}
