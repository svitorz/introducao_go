package enderecos

import "testing"

type cenariosDeTeste struct {
	enderecoRecebido string
	enderecoEsperado string
}

/**
*	comandos para testar:
*	go test
*	go test --cover
* */
func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenariosDeTeste{
		{"Rua Joaquim", "Rua"},
		{"Avenida João Da Costa", "Avenida"},
		{" ", "Tipo Inválido"},
		{"Rodovia Jose Bittencourt", "Rodovia"},
		{"Praça Matriz", "Praça"},
		{"Bairro Centro", "Bairro"},
		{"Cidade Cosmorama", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoRecebido)
		if tipoDeEnderecoRecebido != cenario.enderecoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				tipoDeEnderecoRecebido,
				cenario.enderecoEsperado,
			)
		}
	}
}
