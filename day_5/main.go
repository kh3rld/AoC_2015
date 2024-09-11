package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data []string

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	validCount := SortStr2(data)
	fmt.Println(validCount)
}

func SortStr(str []string) int {
	count := 0
	for _, s := range str {
		if Vow3ls(s) && Double(s) && Ban(s) {
			count++
		}
	}
	return count
}

func Vow3ls(s string) bool {
	vowels := "aeiou"
	count := 0
	for _, v := range s {
		if strings.ContainsRune(vowels, v) {
			count++
		}
	}
	return count >= 3
}

func Double(s string) bool {
	for i := range s {
		if i > 0 && s[i] == s[i-1] {
			return true
		}
	}
	return false
}

func Ban(s string) bool {
	banned := []string{"ab", "cd", "pq", "xy"}
	for _, b := range banned {
		if strings.Contains(s, b) {
			return false
		}
	}
	return true
}

func SortStr2(str []string) int {
	count := 0
	for _, s := range str {
		if UniPaire(s) && B3tween(s) {
			count++
		}
	}
	return count
}

func B3tween(s string) bool {
	for i := 0; i < len(s)-2; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}

func UniPaire(s string) bool {
	m := make(map[string]int)
	for i := 0; i < len(s)-1; i++ {
		p := s[i : i+2]
		if _, exists := m[p]; exists {
			if i > m[p]+1 {
				return true
			}
		} else {
			m[p] = i
		}
	}
	return false
}
