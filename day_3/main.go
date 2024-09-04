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
	dd := Delivery2(string(data))
	fmt.Println(d)
	fmt.Println(dd)
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

func Delivery2(s string) int {
	no := '^'
	so := 'v'
	es := '>'
	we := '<'
	visited := make(map[[2]int]bool)
	count := 0

	santaPos := [2]int{0, 0}
	roboPos := [2]int{0, 0}
	visited[santaPos] = true
	count++

	for i, direction := range s {
		var currentPos *[2]int
		if i%2 == 0 {
			currentPos = &santaPos
		} else {
			currentPos = &roboPos
		}

		switch direction {
		case no:
			currentPos[1]++
		case es:
			currentPos[0]++
		case so:
			currentPos[1]--
		case we:
			currentPos[0]--
		}

		if !visited[*currentPos] {
			visited[*currentPos] = true
			count++
		}
	}
	return count
}
