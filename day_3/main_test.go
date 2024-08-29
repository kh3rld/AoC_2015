package main

import "testing"

func TestDelivery(t *testing.T) {
	got := Delivery("^v^v^v^v^v")
	want := 2
	if want != got {
		t.Errorf("Expected: %v, Got: %v", want, got)
	}
}
