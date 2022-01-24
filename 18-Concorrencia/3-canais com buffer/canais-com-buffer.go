package main

import "fmt"

func main() {
	canal := make(chan string, 2)
	canal <- "Opa gente"

	mensagem := <-canal
	fmt.Println(mensagem)
}
