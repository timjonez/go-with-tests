package main

import "fmt"

const engHello = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return engHello + name
}

func main() {
	fmt.Println(Hello("World"))
}
