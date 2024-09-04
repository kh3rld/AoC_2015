package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	key := "iwrupvqb"
	end := 6
	result := StockingStuffer(key, end)
	fmt.Printf("The answer n is: %d\n", result)
}

func StockingStuffer(key string, end int) int {
	pre := strings.Repeat("0", end)
	n := 1
	for {
		str := key + strconv.Itoa(n)
		hash := md5.Sum([]byte(str))
		hashHex := hex.EncodeToString(hash[:])
		if strings.HasPrefix(hashHex, pre) {
			return n
		}
		n++
	}
}
