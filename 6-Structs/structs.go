package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo Structs")
	var u usuario
	u.nome = "Osvaldo"
	u.idade = 58
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua da Mooca", 3}

	usuario2 := usuario{"Arthur", 28, enderecoExemplo}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 88}
	fmt.Println(usuario3)

}
