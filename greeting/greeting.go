package greeting

import (
	"errors"
	"fmt"
)

const (
	englishHelloPrefix  = "Hello"
	russianHelloPrefix  = "привет"
	japaneseHelloPrefix = "こんにちは"
)

func greetingPrefix(language string) (string, error) {
	switch language {
	case "en":
		return englishHelloPrefix, nil
	case "ru":
		return russianHelloPrefix, nil
	case "jp":
		return japaneseHelloPrefix, nil
	default:
		return "", errors.New("Unsupported language")
	}
}

// Takes the ISO 639-1 code of a supported language as the first argument, and returns a greeting string using the string given as the second argument.
// If an empty string is provided, it uses the string "world" as a placeholder.
func Greet(language, name string) string {
	if name == "" {
		name = "world"
	}

	prefix, err := greetingPrefix(language)

	if err != nil {
		return fmt.Sprintf("%s", err)
	}

	return fmt.Sprintf("%s, %s!", prefix, name)
}
