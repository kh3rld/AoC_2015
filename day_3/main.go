package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	d := Delivery(string(data))
	fmt.Println(d)
}

func Delivery(s string) int {
	no := '^'
	so := 'v'
	es := '>'
	we := '<'
	visited := make(map[[2]int]bool)
	count := 1
	position := [2]int{0, 0}
	visited[position] = true
	for _, direction := range s {
		switch direction {
		case no:
			position[1]++
		case es:
			position[0]++
		case so:
			position[1]--
		case we:
			position[0]--
		}
		if !visited[position] {
			count++
			visited[position] = true
		}
	}
	return count
}
