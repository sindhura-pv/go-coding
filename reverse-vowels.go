package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {

	var l, r int
	l = 0
	r = len(s) - 1
	stringList := strings.Split(s, "")

	for l < r {

		if !isVowel(stringList[l]) {
			l++
		}

		if !isVowel(stringList[r]) {
			r--
		}

		if isVowel(stringList[l]) && isVowel(stringList[r]) {
			temp := stringList[l]
			stringList[l] = stringList[r]
			stringList[r] = temp
			l++
			r--
		}

	}
	return listToString(stringList)

}

func isVowel(s string) bool {

	vowels := [5]string{"a", "e", "i", "o", "u"}
	for _, char := range vowels {
		if strings.ToLower(s) == char {
			return true
		}
	}

	return false

}

func listToString(arr []string) string {

	var s string
	for _, char := range arr {
		s += string(char)
	}
	return s
}

func main() {

	fmt.Print(reverseVowels("hello"))

}
