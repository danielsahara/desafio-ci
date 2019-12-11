package main

import "testing"

func TestSoma5mais5(t *testing.T){
	got := soma(5, 5)
	want := 10

	if got != want{
		t.Errorf("soma(5, 5) got: %v want: %v", got, want)
	}
}

func TestSoma10mais5(t *testing.T){
	got := soma(10, 5)
	want := 15

	if got != want{
		t.Errorf("soma(10, 5) got: %v want: %v", got, want)
	}
}