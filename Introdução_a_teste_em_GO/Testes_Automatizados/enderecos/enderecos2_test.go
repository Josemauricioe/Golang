// Teste de Unidade
package enderecos

import "testing"

type cenarioDeTeste struct{
	enderecoInserido string
	retornoEsperado string
}

func TestTipoDeEndereco2(t *testing.T){
	t.Parallel()
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia br", "Rodovia"},
		{"rua das rosas", "Rua"},
		{"Estrada vaefa", "Estrada"},
		{"abc", "Rua"},
	}

	for _, cenario := range cenariosDeTeste{
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo de endereco %s recebindo é diferente do esperado %s",cenario.enderecoInserido,cenario.retornoEsperado,)


			
		}
	}
}