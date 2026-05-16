package main

import "testing"

func TestHello(t *testing.T) {

	verifyCorrectMessage := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		result := Hello("Jamile")
		expected := "Hello, Jamile"

		verifyCorrectMessage(t, result, expected)
	})

	t.Run("say 'Hello, World' when an empty string is sent", func(t *testing.T) {
		result := Hello("")
		expected := "Hello, World"

		verifyCorrectMessage(t, result, expected)
	})
}
