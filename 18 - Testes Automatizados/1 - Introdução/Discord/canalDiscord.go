package canalDiscord

import "strings"

//TipoDoCanal retorna se é valido a funcionalidade do canal do Discord (Sua função principal)
func TipoDoCanal(tipo string) string {
	tiposValidos := []string{"Chat", "Log", "Voip", "Eventos"}

	palavraTipoMinusculo := strings.ToLower(tipo)
	palavraTipo := strings.Split(palavraTipoMinusculo, " ")[0]

	tipoEValido := false
	for _, tipo:= range tiposValidos{
		if tipo ==palavraTipo{
			tipoEValido = true
		}
	}
	
	if tipoEValido {
		return palavraTipo
	}
	return "Tipo Inválido"
}