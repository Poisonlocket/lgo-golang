package hello

import "fmt"

// Calls The Hello function to print a string
func main() {
	fmt.Println(Hello("Johnny"))
}

const englishHelloPrefix = "Hello, "

// returns a string Hello, World
func Hello(name string) string {
	return englishHelloPrefix + "World!"
}
