package main

import "fmt"

const (
	English    = "english"
	Portuguese = "portuguese"
	Spanish    = "spanish"
	French     = "french"
)

var prefixes = map[string]string{
	English:    "Hello, ",
	Portuguese: "Oi, ",
	Spanish:    "Hola, ",
	French:     "Bonjour, ",
}

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix, ok := prefixes[lang]
	if !ok {
		prefix = "Hello, " // default value
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}
