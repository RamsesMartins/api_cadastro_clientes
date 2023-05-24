package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/ramses/api-clientes/config"
)

func OppeningDatabase() (*sql.DB, error) {
	cfg, _, err := config.GetConfigs()
	if err != nil {
		fmt.Println(err.Error())
	}

	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s", cfg.User, cfg.Password, cfg.Hostname, cfg.Port, cfg.Database))
	if err != nil {
		fmt.Println(err.Error())
	}
	return conn, err
}
