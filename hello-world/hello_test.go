package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello()
	expected := "Hello, world!"
	
	if result != expected {
		t.Errorf("Result '%s', expected '%s'", result, expected)
	}
}