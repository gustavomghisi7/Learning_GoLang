package main

import (
	"fmt"
	"time"
)

func main() {
	// concorrencia != paralelismo
	go escrever("Ola Mundo!!!") //goroutine
	escrever("Programando em GoLang!!!!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
