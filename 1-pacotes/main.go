package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo no arquivo main")
	auxiliar.Escrever()
	// auxiliar.escrever2()

	erro := checkmail.ValidateFormat("denmiran.gmail.com")
	fmt.Println(erro)
}
