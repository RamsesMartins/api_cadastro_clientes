package models

import (
	"fmt"
	"github.com/ramses/api-clientes/database"
)

func GetCliente(id int) (Clientes, error) {
	var cliente Clientes
	db, err := database.OppeningDatabase()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	row := db.QueryRow("SELECT cliente_id,nome, sobrenome, dtnascimento FROM clientes WHERE cliente_id = ?", id)
	row.Scan(&cliente.Cliente_id, &cliente.Nome, &cliente.Sobrenome, &cliente.Dtnascimento)
	return cliente, nil
}
