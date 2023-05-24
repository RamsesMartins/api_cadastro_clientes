package models

import (
	"fmt"

	"github.com/ramses/api-clientes/database"
)

//ROTA PARA RETORNAR APENAS UM CLIENTE
func InsertClientes(clientes Clientes) (Clientes, error) {
	database, err := database.OppeningDatabase()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer database.Close()
	_, err = database.Exec("INSERT INTO clientes(nome, sobrenome, dtnascimento) VALUES(?, ?, ?)", clientes.Nome, clientes.Sobrenome, clientes.Dtnascimento)
	if err != nil{
		fmt.Println(err.Error())
	}

	return clientes, nil
}
