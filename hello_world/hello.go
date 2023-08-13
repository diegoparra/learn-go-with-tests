package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	englishHelloPrefix = "Hello, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return

}

func main() {
	fmt.Println(Hello("parra", ""))
}
