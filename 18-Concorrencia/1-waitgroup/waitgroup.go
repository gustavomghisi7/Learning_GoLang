package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitgroup sync.WaitGroup

	waitgroup.Add(4)

	go func() {
		escrever("Ola Mundo!!!")
		waitgroup.Done() // -1
	}()

	go func() {
		escrever("Programando em GoLang!!!!")
		waitgroup.Done() // -1
	}()

	go func() {
		escrever("Goroutine 3")
		waitgroup.Done() // -1
	}()

	go func() {
		escrever("Goroutine 4")
		waitgroup.Done() // -1
	}()

	waitgroup.Wait() // 2
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
