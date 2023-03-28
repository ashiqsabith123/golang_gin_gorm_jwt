package connectdb

import (
	"fmt"
	"os"

	error "golang_gin_gorm_jwt/checkError"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {

	conn_str := os.Getenv("CONN_STR")

	db, err := gorm.Open(postgres.Open(conn_str), &gorm.Config{})
	error.CheckError(err)

	fmt.Println("Connected suceesfully....!")

	return db
}
