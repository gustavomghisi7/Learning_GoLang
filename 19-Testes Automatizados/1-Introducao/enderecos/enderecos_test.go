package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		// {"Praca das Rosas", "Tipo Invalido"},
		{"Estrada da Lagrimas", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUCAS", "Avenida"},
		// {" ", "Tipo Invalido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s eh diferente do esperado %s!!!  ",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}

	// enderecoParaTeste := "Avenida ABC"
	// tipoDeEnderecoEsperado := "Avenida"
	// tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	// if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
	// 	t.Errorf("O tipo recebido eh diferente do esperado!!! Esperado %s  Recebido %s",
	// 		tipoDeEnderecoEsperado,
	// 		tipoDeEnderecoRecebido)
	// }
}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Error("Teste quebrou")
	}
}
