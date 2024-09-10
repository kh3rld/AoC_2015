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

	validCount := SortStr(data)
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
