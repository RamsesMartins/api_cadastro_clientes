# Cadastro de Clientes
---
Tecnologias | Função
------------ | -------------
| <img src="https://img.shields.io/badge/MariaDB-01529E?style=for-the-badge&logo=mariadb&logoColor=white" /> | `Banco de Dados` |
| <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" /> | `Linguagem` |
| <img src="https://img.shields.io/badge/Windows-017AD7?style=for-the-badge&logo=windows&logoColor=white" /> | `Sistema Operacional` |
--------
# Como instalar?
## Requisitos
- Git
- Go Compiler
- Maria DB / MySQL

## Configurando
Para começar, é preciso que seja clonado este repositório em algum diretório seu. Caso não saiba como clonar um repositório hospedado no GitHub veja este [link](https://docs.github.com/pt/repositories/creating-and-managing-repositories/cloning-a-repository).

Em seguida é nescessário instalar o banco de dados Maria DB, ou se preferir o MySQL. Após a instalação acesse o seu terminal e acesse o utilitário de linha de comando do SGBD para podermos escrever nossos códigos SQL.

Será preciso criar um Database. No exemplo abaixo é ilustrado o comando para executar tal ação:
````sql
CREATE DATABASE IF NOT EXISTS <NOME_DATABASE>
````
Importante alterar `<NOME_DATABASE>` para o nome que você realmente deseja colocar.

Após a criação do datbase, podemos agora criar tabela `cliente` que irá conter os dados dos clientes cadastrados pela API.
```sql
CREATE TABLE clientes (
  cliente_id int(11) NOT NULL AUTO_INCREMENT,
  nome varchar(25) NOT NULL,
  sobrenome varchar(50) NOT NULL,
  dtnascimento date NOT NULL,
  PRIMARY KEY (`cliente_id`)
)
```
**Pronto!!!** Agora você já tem o Banco de Dados Configurado. Precisamos apenas passar as informações para a aplicação em Go.

Para coeçarmos devemos ir para dentro do diretório clonado. Veja se a estrutura de pastas está assim:
```
api_cliente -->
                config -->
                conrollers -->
                database -->
                models -->
                go.mod
                go.sum
                main.go

```
Dentro deste diretório (`api_cliente`) você precisará criar o arquivo **_.env_**, é neste arquivo que iremos passar os parâmetros de conexão do nosso banco de dados

Abra o arquivo **_.env_** e escreva o seguinte:
```yaml
#API CONFIG
API_PORT = "3030"
#DATABASE CONFIG
DB_HOSTNAME = "localhost"
DB_PORT = "3306"
DB_USER = "root"
DB_PASSWORD = "senhadousuario"
DB_DATABASE = "<NOME_DATABASE>"

```
Edite os parâmetros acima de acordo com o seu desejo.

Após a configuração do **_.env_** sua estrutura deve ficar assim:
```
api_cliente -->
                config -->
                conrollers -->
                database -->
                models -->
                go.mod
                go.sum
                main.go
                .env

```

# Endpoints
Método | Rotas | Função
--- | ------ | ---
| **`GET`** | /clientes |Trará como resultado um array json com todos os clientes cadastrados|
| **`GET`** | /clientes?id=x |Trará como resultado o _JSON_ com dados do cliente que tiver o cliente_id semelhante ao id passado|
| **`POST`** | /clientes |Trará como resultado o _JSON_ com dados do cliente for cadastrado|
| **`PUT`** | /clientes?id=x |Modificará os dados do cliente seguindo os dados passados no body _JSON_|
| **`DELETE`** | /clientes?id=x |Apagará o cliente do banco de dados|
