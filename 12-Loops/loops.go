package main

import (
	"fmt"
	"time"
)

func main() {
	// i := 0
	// for i < 10 {
	// 	i++
	// 	fmt.Println("incrementando i")
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)

	// for j := 0; j < 10; j += 2 {
	// 	fmt.Println("incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	// nomes := []string{"Joao", "Davi", "Lucas"}

	// for _, valor := range nomes {
	// 	fmt.Println(valor)
	// }

	// for indice, letra := range "PALAVRA" {
	// 	fmt.Println(indice, string(letra))
	// }

	// usuario := map[string]string{
	// 	"nome":      "Leonardo",
	// 	"sobrenome": "Silva",
	// }

	// for chave, valor := range usuario {
	// 	fmt.Println(chave, valor)
	// }

	// type usuarioStruct struct {
	// 	nome string
	// 	sobrenome string
	// }

	// usuario2 := usuarioStruct("Osvaldo", "Silva")
	// for chave, valor := range usuarioStruct {
	// 	fmt.Println(chave, valor)
	// }

	for true {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}

}
