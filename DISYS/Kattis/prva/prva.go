package main

import (
	"fmt"
)

func main() {
	var R, C int
	fmt.Scanln(&R, &C)
	var chars = make([][]string, R, C)
	for i := 0; i < R; i++ {
		var word string
		fmt.Scan(&word)
		for j := 0; j < C; j++ {
			chars[i] = append(chars[j], string(word[j]))
		}
	}
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			fmt.Print(chars[j])
		}
	}

}
