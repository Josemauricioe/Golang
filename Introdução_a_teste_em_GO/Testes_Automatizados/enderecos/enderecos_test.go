// Teste de Unidade
package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T){
	t.Parallel()
	enderecoParaTeste := "Avenida Paulista"
	tipoDeEnderecoEsperado := "Avenida"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)
 
	if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
		t.Error("O tipo recebindo é diferente do esperado!!!")
	}
}

func TestQualquer(t *testing.T){
	t.Parallel()
	if 1 > 2 {
		t.Error("Teste quebrou!")
	}
}