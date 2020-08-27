package main

import "fmt"

const str string = "Hello, "
const spanish string = "Spanish"
const strSpanish string = "Hola, "
const French string = "French"
const strFrench string = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		return greetingPrefix(language) + "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = strSpanish
	case French:
		prefix = strFrench
	default:
		prefix = str
	}
	return
}

func main() {
	fmt.Println("hello world")
}
