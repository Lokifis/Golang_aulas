package canalDiscord

import "testing"

type cenarioDeTeste struct {
	canalInserido   string
	retornoEsperado string
}

func TestTipoDeCanal(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Chat-geral", "Chat"},
		{"Chat-geral", "Chat"},
		{"Chat-geral", "Chat"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDoCanal(cenario.canalInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("Tipo %s diferente do esperado %s", retornoRecebido, cenario.retornoEsperado)
		}
	}
}
