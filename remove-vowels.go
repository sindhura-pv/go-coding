package main

func removeVowels(S string) string {

	var res string

	for _, char := range S {

		if !isvowel(string(char)) {
			res += string(char)
		}

	}

	return res

}

func isvowel(s string) bool {

	vowels := [5]string{"a", "e", "i", "o", "u"}
	for _, char := range vowels {
		if s == char {
			return true
		}
	}

	return false

}
