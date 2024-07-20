package main

import (
	"fmt"
	"log"
	"net/http"

	dependencyInjection "github.com/maeveblackmore/learn-go-with-tests/dependency-injection"
	"github.com/maeveblackmore/learn-go-with-tests/greeting"
)

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	dependencyInjection.Greet(w, "world")
}

func main() {
	prefixLanguage := "jp"
	firstName := "Maeve"
	fmt.Println(greeting.Greet(prefixLanguage, firstName))
	port := 5001
	address := fmt.Sprintf(":%d", port)

	log.Fatal(http.ListenAndServe(address, http.HandlerFunc(MyGreetHandler)))
}
