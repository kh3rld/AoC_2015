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

func TestSortStr2(t *testing.T) {
	s := []string{"uurcxstgmygtbstg", "ieodomkazucvgmuy", "qjhvhtzxzqqjkmpb"}
	want := 1
	got := SortStr2(s)
	if got != want {
		t.Errorf("Expected %v, Got: %v", want, got)
	}
}
