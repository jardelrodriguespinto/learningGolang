package model

import "time"

type Pessoa struct {
	Nome             string
	Idade            int
	Endereco         Endereco
	DataDeNascimento time.Time
}

func (p *Pessoa) ModificaNome(nome string) {
	p.Nome = nome
}

func (p Pessoa) IdadeAtual() int {
	anoDeNascimento := p.DataDeNascimento.YearDay()
	anoAtual := time.Now().Year()
	return anoAtual - anoDeNascimento
}

func CalculaIdade(p Pessoa) int {
	anoDeNascimento := p.DataDeNascimento.YearDay()
	anoAtual := time.Now().Year()
	return anoAtual - anoDeNascimento
}
