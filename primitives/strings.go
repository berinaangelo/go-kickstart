package main

import (
	"fmt"
	"slices"
)

func initStrings() {
	string1 := "Hello There"

	// substrings
	fmt.Println(string1[3:5]) // lo
	fmt.Println(string1[:5])

	// runes - numeric representation of characters in string
	runes := []rune(string1)
	fmt.Println(runes)

	// casting it back to string converts it back to characters
	fmt.Println(string(runes))

	// reversing a string
	slices.Reverse(runes)
	fmt.Println(string(runes))
}
