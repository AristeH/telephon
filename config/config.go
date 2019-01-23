package config

import (
	"database/sql"

//	_ "github.com/nakagami/firebirdsql"
)

// Config  Таблица конфигурации. Необходимые настрройки для приложения.
type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
	DB   *sql.DB
}

// Parametrs  создание подключения к БД
var Parametrs Config

// InitDB  создание подключения к БД
func setDB (db *sql.DB) {
	Parametrs.DB = db
}


