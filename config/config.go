package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Substitua as informações de conexão do PostgreSQL abaixo
	db, err = gorm.Open("postgres", "host=localhost port=32768 user=postgres dbname=postgres password=postgrespw sslmode=disable")

	if err != nil {
		return err
	}

	// Outras configurações do GORM, como definir um pool de conexão ou modo de log, podem ser feitas aqui

	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
