package greeting

import "fmt"

const (
	englishHelloPrefix  = "Hello"
	russianHelloPrefix  = "привет"
	japaneseHelloPrefix = "こんにちは"
)

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "ru":
		prefix = russianHelloPrefix
	case "jp":
		prefix = japaneseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

// Takes the ISO 639-1 code of a language as the first argument, and returns a greeting string using the string given as the second argument.
// If an empty string is provided, it uses the string "world" as a placeholder.
func Greet(language, name string) string {
	if name == "" {
		name = "world"
	}

	return fmt.Sprintf("%s, %s!", greetingPrefix(language), name)
}
