package main

import (
	"fmt"
	"hello/greeting"
	"hello/integers"
)

func main() {
	prefixLanguage := "jp"
	firstName := "Maeve"
	fmt.Println(greeting.Greet(prefixLanguage, firstName))

	integers.Add(1, 2)
}
