package models

import (
	"fmt"

	"github.com/ramses/api-clientes/database"
)

func GetAll() (clientes []Clientes, err error) {
	db, err := database.OppeningDatabase()
	if err != nil {
		fmt.Println(err.Error())
	}
	rows, err := db.Query(`SELECT cliente_id, nome, sobrenome, dtnascimento FROM clientes`)
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		var cliente Clientes
		err := rows.Scan(&cliente.Cliente_id, &cliente.Nome, &cliente.Sobrenome, &cliente.Dtnascimento)
		if err != nil {
			fmt.Println(err.Error())
		}
		clientes = append(clientes, cliente)
	}
	return clientes, nil

}
