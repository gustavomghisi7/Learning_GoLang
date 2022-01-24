package main

import "fmt"

func main() {
	var variavel1 string = "variavel1"
	variavel2 := "variavel2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "variavel3"
		variavel4 string = "variavel4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "variavel5", "variavel6"
	fmt.Println(variavel5, variavel6)

	const constante1 string = "constante 1"
	fmt.Println(constante1)

	// trocando valores das variaveis
	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)

}
