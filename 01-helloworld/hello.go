package helloworld

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}
	greetingPrefix := englishHelloPrefix

	switch lang {
	case spanish:
		greetingPrefix = spanishHelloPrefix
	case french:
		greetingPrefix = frenchHelloPrefix
	}

	return greetingPrefix + name
}
