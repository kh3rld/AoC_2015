package main

import "testing"

func TestMain(t *testing.T) {
	want := SurfaceArea("2x3x4")
	got := 58
	if want != got {
		t.Errorf("Expected %v, Got %v", want, got)
	}
}
