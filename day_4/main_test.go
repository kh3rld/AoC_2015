package main

import (
	"testing"
)

func TestStockingStuffer(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"abcdef", 609043},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StockingStuffer(tt.name, 5); got != tt.want {
				t.Errorf("StockingStuffer() = %v, want %v", got, tt.want)
			}
		})
	}
}
