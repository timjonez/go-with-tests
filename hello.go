package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const russian = "Russian"

const ruHello = "Привет, "
const frHello = "Bonjour, "
const espHello = "Hola, "
const engHello = "Hello, "

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := engHello
	switch lang {
	case spanish:
		prefix = espHello
	case french:
		prefix = frHello
	case russian:
		prefix = ruHello
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("World", ""))
}
