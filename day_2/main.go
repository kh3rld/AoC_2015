package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	s := strings.Split(string(data), "\n")
	var result int
	var result2 int
	for _, x := range s {
		result += SurfaceArea(x)
		result2 += SurfaceArea2(x)
	}
	fmt.Println(result)
	fmt.Println(result2)
}

func SurfaceArea(s string) int {
	ss := strings.Split(s, "x")
	l, _ := strconv.Atoi(ss[0])
	w, _ := strconv.Atoi(ss[1])
	h, _ := strconv.Atoi(ss[2])
	length := 2 * (l * w)
	width := 2 * (w * h)
	height := 2 * (h * l)
	slack := (min(length, min(width, height))) / 2
	r := length + width + height
	return r + slack
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func SurfaceArea2(s string) int {
	ss := strings.Split(s, "x")
	l, _ := strconv.Atoi(ss[0])
	w, _ := strconv.Atoi(ss[1])
	h, _ := strconv.Atoi(ss[2])
	length := 2 * (l + w)
	width := 2 * (w + h)
	height := 2 * (h + l)
	slack := (min(length, min(width, height)))
	return slack + l*w*h
}
