package main

import (
	"fmt"
	"os"

	dependencyInjection "github.com/maeveblackmore/learn-go-with-tests/dependency-injection"
	"github.com/maeveblackmore/learn-go-with-tests/greeting"
	"github.com/maeveblackmore/learn-go-with-tests/integers"
)

func main() {
	prefixLanguage := "jp"
	firstName := "Maeve"
	fmt.Println(greeting.Greet(prefixLanguage, firstName))
	dependencyInjection.Greet(os.Stdout, firstName)
	integers.Add(1, 2)
}
