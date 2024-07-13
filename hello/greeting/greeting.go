package greeting

import "fmt"

const (
	englishHelloPrefix  = "Hello"
	russianHelloPrefix  = "привет"
	japaneseHelloPrefix = "こんにちは"
)

func Greet(prefix, name string) string {
	if name == "" {
		name = "world"
	}

	switch prefix {
	case "en":
		prefix = englishHelloPrefix
	case "ru":
		prefix = russianHelloPrefix
	case "jp":
		prefix = japaneseHelloPrefix
	}

	return fmt.Sprintf("%s, %s!", prefix, name)
}
