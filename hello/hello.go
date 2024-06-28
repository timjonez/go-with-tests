package main

import "fmt"

const (
	spanish = "Spanish"
	french  = "French"
	russian = "Russian"

	ruHello  = "Привет, "
	frHello  = "Bonjour, "
	espHello = "Hola, "
	engHello = "Hello, "
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	return getGreeting(lang) + name
}

func getGreeting(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frHello
	case spanish:
		prefix = espHello
	case russian:
		prefix = ruHello
	default:
		prefix = engHello
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
