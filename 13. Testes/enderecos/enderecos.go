package enderecos

import "strings"

/**
*	@param Um endereço qualquer
*	@return Se o endereço está entre os validados pelo sistema.
* */
func TipoDeEndereco(endereco string) string {
	tiposAceitos := [5]string{"rua", "avenida", "rodovia", "praça", "bairro"}

	endereco = strings.Split(strings.ToLower(endereco), " ")[0]

	for _, tipo := range tiposAceitos {
		if tipo == endereco {
			return strings.Title(endereco)
		}
	}
	return "Tipo Inválido"
}
