package main

import "fmt"

const prefixHelloPortuguese = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return prefixHelloPortuguese + name
}

func main() {
	fmt.Println(Hello(""))
}
