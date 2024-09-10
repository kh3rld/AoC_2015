package main

import (
	"testing"
)

func TestSortStr(t *testing.T) {
	s := []string{"ugknbfddgicrmopn", "jchzalrnumimnmhp", "haegwjzuvuyypxyu", "dvszwmarrgswjxmb", "ugknbfddgicrmopn"}
	want := 2
	got := SortStr(s)
	if got != want {
		t.Errorf("Expected %v, Got: %v", want, got)
	}
}
