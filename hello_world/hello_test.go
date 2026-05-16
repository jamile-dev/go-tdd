package main

import "testing"

func TestHello(t *testing.T) {
	result := Hello("Jamile")
	expected := "Hello, Jamile"

	if result != expected {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}
