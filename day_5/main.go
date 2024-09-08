package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, _ := os.ReadFile("input.txt")
	s := strings.Split(string(data), "\n")
	c := 0
	for _, x := range s {
		if SortStr(x) {
			c++
		}
	}
	fmt.Println(c)
}

func SortStr(s string) bool {
	vow := "aeiouAEIOU"
	ban := []string{"ab", "cd", "pq", "xy"}
	count := 0
	for i := range s {
		for _, v := range vow {
			for _, b := range ban {
				if strings.ContainsRune(s, v) {
					count++
				}
				if count >= 3 {
					if i > 0 && s[i] == s[i-1] {
						if strings.Contains(s, b) {
							return true
						}
					}
				}
			}
		}
	}
	return false
}
