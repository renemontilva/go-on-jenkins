package main

import "testing"

func MainTest(t *testing.T) {
	name := "Hello World"
	got := Hello(name)
	want := "Hello World"
	if  got != want {
		t.Error("Error")
	}
}
