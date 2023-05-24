package models

type Clientes struct {
	Cliente_id   int64  `json:"cliente_id"`
	Nome         string `json:"nome"`
	Sobrenome    string `json:"sobrenome"`
	Dtnascimento string `json:"dtnascimento"`
}
