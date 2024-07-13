package main

import (
	"fmt"
	"hello/greeting"
)

func main() {
	prefixLanguage := "jp"
	firstName := "Maeve"
	fmt.Println(greeting.Greet(prefixLanguage, firstName))
}
