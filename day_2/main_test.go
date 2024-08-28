package main

import "testing"

func TestSurfaceArea(t *testing.T) {
	want := SurfaceArea("2x3x4")
	got := 58
	if want != got {
		t.Errorf("Expected %v, Got %v", want, got)
	}
}

func TestSurfaceArea2(t *testing.T) {
	want := SurfaceArea2("1x1x10")
	got := 14
	if want != got {
		t.Errorf("Expected %v, Got %v", want, got)
	}
}
