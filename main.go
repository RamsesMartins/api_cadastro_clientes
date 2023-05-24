package main

import (
	"fmt"
	"net/http"

	"github.com/ramses/api-clientes/config"
	controller "github.com/ramses/api-clientes/controllers"
)

func main() {
	db, api, err := config.GetConfigs()
	if err != nil {
		panic(err.Error())
	}
	//FEEDBACK
	fmt.Println("-----SERVIDOR RODANDO-----")
	fmt.Printf("|Porta da API: %s    \n", api.Port)
	fmt.Printf("|Banco de Dados: %s  \n", db.Database)
	fmt.Printf("|Usuário BD: %s      \n", db.User)
	fmt.Println("--------------------------")
	fmt.Println("\nCriado por Ramsés de Oliveira Martins")

	// ENDPOINTS
	http.HandleFunc("/", controller.GET_IndexPage)
	http.HandleFunc("/clientes", controller.Controller_Cliente)
	// LISTENING
	http.ListenAndServe(fmt.Sprintf(":%s",api.Port) ,nil)

	
}
