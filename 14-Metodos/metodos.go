package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("salvando os dados do usuario %s no banco de dados\n", u.nome)
}

func main() {
	usuario1 := usuario{"Osvaldo", 20}
	// fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Vera", 48}
	usuario2.salvar()

}
