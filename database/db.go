package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB     *gorm.DB
	err    error
	use_db int // 1 for using DB, 0 for using mock (default as not initialized)
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		//log.Panic("Erro ao conectar com banco de dados")
		fmt.Println("Erro ao conectar com banco de dados")
	} else {
		use_db = 1
	}
}

func UseDB() int {
	return use_db
}
