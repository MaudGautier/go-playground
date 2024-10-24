package main

import (
	"fmt"
	"strings"
)

func ParseString(s string) {
	fmt.Println("--- Parsing per rune:")
	for _, letter := range s {
		fmt.Printf("Type: %T, Value: %v, Char: %c\n", letter, letter, letter)
	}

	fmt.Println("\n--- Parsing per word")
	for _, word := range strings.Split(s, " ") {
		fmt.Printf("Type: %T, Value: %v, Char: %c\n", word, word, word)
	}
}

func main() {
	ParseString("I am learning Go!")
	/*
		--- Parsing per rune:
		Type: int32, Value: 73, Char: I
		Type: int32, Value: 32, Char:
		Type: int32, Value: 97, Char: a
		Type: int32, Value: 109, Char: m
		Type: int32, Value: 32, Char:
		Type: int32, Value: 108, Char: l
		Type: int32, Value: 101, Char: e
		Type: int32, Value: 97, Char: a
		Type: int32, Value: 114, Char: r
		Type: int32, Value: 110, Char: n
		Type: int32, Value: 105, Char: i
		Type: int32, Value: 110, Char: n
		Type: int32, Value: 103, Char: g
		Type: int32, Value: 32, Char:
		Type: int32, Value: 71, Char: G
		Type: int32, Value: 111, Char: o
		Type: int32, Value: 33, Char: !

		--- Parsing per word:
		Type: string, Value: I, Char: %!c(string=I)
		Type: string, Value: am, Char: %!c(string=am)
		Type: string, Value: learning, Char: %!c(string=learning)
		Type: string, Value: Go!, Char: %!c(string=Go!)
	*/
}

/*
Explanations:
- `range s` when `s` is a string ranges over the Unicode code points of the characters in the string,
  not their ASCII values. In Go, the Unicode code points are returned as a rune (i.e. an alias for int32)
  => the values correspond to the Unicode point value for the letter.
- `strings.Split` returns an array of strings => the values correspond to each word letters
*/
