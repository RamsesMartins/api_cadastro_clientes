package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/ramses/api-clientes/models"
)

func Controller_Cliente(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	if r.Method == http.MethodGet{		
		if id != "" {
			GET_getCliente(w,r,id)
		}else{
			GET_getAll(w,r)
		}

	}else if r.Method == http.MethodPost{
		POST_insertCliente(w,r)

	}else if r.Method == http.MethodPut{
		PUT_updateCliente(w,r, id)
	}else if r.Method == http.MethodDelete{
		DELETE_cliente(w,r,id)
	}else{
		http.Error(w,"Método não permitido", http.StatusMethodNotAllowed)
	}
}

// ROTA PADRÃO
func GET_IndexPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "pages/index.html")
}

// FUNÇÃO PARA RETORNAR TODOS OS CLIENTES
func GET_getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var clientes []models.Clientes
	clientes, err := models.GetAll()
	if err != nil {
		http.Error(w, "Erro ao buscar clientes cadastrados", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(clientes)
}

// FUNÇÃO PARA CADASTRAR UM CLIENTE
func POST_insertCliente(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error ao ler o body da requisição", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var cliente models.Clientes
	err = json.Unmarshal(body, &cliente)
	if err != nil {
		http.Error(w, "Error ao decodificar JSON", http.StatusInternalServerError)
		return
	}
	models.InsertClientes(cliente)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cliente)
}

// FUNÇÃO PARA RETORNAR APENAS UM CLIENTE
func GET_getCliente(w http.ResponseWriter, r *http.Request, id string) {
	cliente_id, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cliente, err := models.GetCliente(cliente_id)
	if err != nil {
		http.Error(w, "Erro ao buscar cliente neste id", http.StatusInternalServerError)
		return
	}
	if cliente.Cliente_id == 0{
		http.Error(w, "Erro ao buscar cliente", http.StatusNoContent)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cliente)

}

// FUNÇÃO PARA ATUALIZAR A INFORMAÇÃO DE UM CLIENTE
func PUT_updateCliente(w http.ResponseWriter, r *http.Request, cliente_id string) {
	if r.Method != http.MethodPut {
		http.Error(w, "Método não permitido nesta rota", http.StatusMethodNotAllowed)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := strconv.Atoi(cliente_id)
	var cliente models.Clientes
	err = json.Unmarshal(body, &cliente)
	if err != nil {
		http.Error(w, "Erro ao decodificar json", http.StatusInternalServerError)
		return
	}

	res, err := models.UpdateCliente(cliente, id)
	if err != nil {
		http.Error(w, "Erro ao atualizar o cliente", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

//FUNÇÃO PARA DELETAR UM CLIENTE
func DELETE_cliente(w http.ResponseWriter, r *http.Request, cliente_id string){
	id, err := strconv.Atoi(cliente_id)
	if err != nil {
		http.Error(w, "Erro ao converter ID de string para int", http.StatusInternalServerError)
		return
	}
	res, err:=models.DeleteCliente(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(res))
}