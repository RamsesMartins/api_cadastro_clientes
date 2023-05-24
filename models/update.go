package models

import (
	"fmt"

	"github.com/ramses/api-clientes/database"
)

func UpdateCliente(cliente Clientes, id int) (Clientes, error) {
	db, err := database.OppeningDatabase()
	if err != nil {
		fmt.Println(err.Error())
		db.Close()
		return cliente, nil
		
	}
	defer db.Close()
	prepare, err := db.Prepare("UPDATE clientes SET nome = ?, sobrenome = ? WHERE cliente_id = ?")
	if err != nil {
		fmt.Print(err.Error())
		return cliente, err
	}

	res, err := prepare.Exec(cliente.Nome, cliente.Sobrenome, id)
	if err != nil {
		fmt.Print(err.Error())
		return cliente, err
	}
	fmt.Print(res)
	return cliente, nil
}