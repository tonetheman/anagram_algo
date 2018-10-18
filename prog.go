package main

import "fmt"

var cton = map[rune]int{
	'a': 2,
	'b': 3,
	'c': 5,
	'd': 7,
	'e': 11,
}

// this will score a word
// you can take the result and compare
// to another words
func score(word string) int {
	sum := 0
	for x, c := range word {
		fmt.Println(x, c, cton[c])
		sum += cton[c]
	}
	return sum
}

func main() {
	fmt.Println(score("abc"))
	fmt.Println(score("aab"))

	fmt.Println(score("abc") == score("bca"))
}
