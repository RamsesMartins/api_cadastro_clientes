package models

import (
	"fmt"

	"github.com/ramses/api-clientes/database"
)

func DeleteCliente(id int) (string, error) {
	db, err := database.OppeningDatabase()
	if err != nil {
		return "erro",err
	}
	defer db.Close()
	prepare, err := db.Prepare("DELETE FROM clientes WHERE cliente_id = ?")
	if err != nil {
		return "erro", err
	}

	res, err := prepare.Exec(id)
	if err != nil {
		return "erro", err
	}
	fmt.Print(res)
	return "Cliente Excluído", nil
}