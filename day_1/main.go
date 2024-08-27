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
	f2 := NotQuiteLisp2(string(input))
	fmt.Printf("Floor: %v", f)
	fmt.Println()
	fmt.Printf("Floor: %v", f2)
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

func NotQuiteLisp2(s string) int {
	f := 0
	r := 0
	for i, x := range s {
		if x == '(' {
			f++
		} else if x == ')' {
			f--
		}
		if f == -1 {
			r = i + 1
			break
		}
	}
	return r
}
