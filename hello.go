package main

import "fmt"

const spanish = "Spanish"
const english = "English"
const french = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	var prefix string

	if name == "" {
		name = "World"
	}

	if language == "" {
		language = english
	}

	if language == english {
		prefix = englishHelloPrefix
	}

	if language == spanish {
		prefix = spanishHelloPrefix
	}

	if language == french {
		prefix = frenchHelloPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("world", "English"))
}
