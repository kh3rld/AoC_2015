package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading Input data")
		return
	}
	f := NotQuiteLisp(string(input))
	fmt.Printf("Floor: %v", f)
	fmt.Println()
}

func NotQuiteLisp(s string) int {
	f := 0
	for _, x := range s {
		if x == '(' {
			f++
		} else if x == ')' {
			f--
		}
	}
	return f
}
