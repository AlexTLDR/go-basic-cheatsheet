package main

import "testing"

func ExampleMain() {
	main()
	// Output:
	// Hello world
}

func TestGreet_English(t *testing.T) {
	lang := "en"
	expectedGreeting := "Hello world"

	greeting := greet(language(lang))

	if greeting != expectedGreeting {
		// mark this test as failed
		t.Errorf("expected: %q, got: %q", expectedGreeting, greeting)
	}
}

func TestGreet_French(t *testing.T) {
	lang := "fr"
	expectedGreeting := "Bonjour le monde"

	greeting := greet(language(lang))

	if greeting != expectedGreeting {
		// mark this test as failed
		t.Errorf("expected: %q, got: %q", expectedGreeting, greeting)
	}
}

func TestGreet_Romanian(t *testing.T) {
	lang := "ro"
	expectedGreeting := ""

	greeting := greet(language(lang))

	if greeting != expectedGreeting {
		// mark this test as failed
		t.Errorf("expected: %q, got: %q", expectedGreeting, greeting)
	}
}
