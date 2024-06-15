package hello

const (
	helloEnglishPrefix = "Hello "
	helloSpanishPrefix = "Hola "
	helloDeutchPrefix  = "Hallo "

	deutch  = "Deutch"
	spanish = "Spanish"

// emptyPhrase = "Hello World !"
)

func Hello(name string, langage string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(langage) + name + " !"
}

func greetingPrefix(langage string) (helloPrefix string) {
	switch langage {
	case spanish:
		helloPrefix = helloSpanishPrefix
	case deutch:
		helloPrefix = helloDeutchPrefix
	default:
		helloPrefix = helloEnglishPrefix
	}
	return
}
