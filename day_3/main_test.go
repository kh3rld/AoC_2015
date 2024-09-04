package main

import (
	"testing"
)

func TestDelivery(t *testing.T) {
	got := Delivery("^v^v^v^v^v")
	want := 2
	if want != got {
		t.Errorf("Expected: %v, Got: %v", want, got)
	}
}

func TestDelivery2(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"^>v<", 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Delivery2(tt.name); got != tt.want {
				t.Errorf("Delivery2() = %v, want %v", got, tt.want)
			}
		})
	}
}
