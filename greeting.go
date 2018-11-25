package hello

const french = "French"
const italian = "Italian"
const portuguese = "Portuguese"
const spanish = "Spanish"

const defaultPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const italianHelloPrefix = "Ciao, "
const portugueseHelloPrefix = "Olá, "
const spanishHelloPrefix = "Hola, "

// Greeting returns a greeting message
func Greeting(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	switch language {
	case french:
		return frenchHelloPrefix
	case italian:
		return italianHelloPrefix
	case portuguese:
		return portugueseHelloPrefix
	case spanish:
		return spanishHelloPrefix
	default:
		return defaultPrefix
	}
}
