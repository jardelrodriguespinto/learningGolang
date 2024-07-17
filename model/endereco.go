package model

import "strconv"

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
}

func (e Endereco) GetEnderecoCompleto(endereco Endereco) string {
	enderecoCompleto := e.Rua + " " + strconv.Itoa(e.Numero) + " " + e.Cidade

	return enderecoCompleto
}
