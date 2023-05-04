package models

type Clients struct {
	ID             int    `json:"id"`
	Nome           string `json:"nome"`
	Sobrenome      string `json:"sobrenome"`
	Contato        string `json:"contato"`
	Endereco       string `json:"endereco"`
	DataNascimento string `json:"data_nascimento"`
	CPF            string `json:"cpf"`
}
