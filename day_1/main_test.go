package main

import "testing"

func TestNotQuiteLisp(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"(())", 0},
		{"(((", 3},
		{"))(((((", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NotQuiteLisp(tt.name); got != tt.want {
				t.Errorf("NotQuiteLisp() = %v, want %v", got, tt.want)
			}
		})
	}
}
