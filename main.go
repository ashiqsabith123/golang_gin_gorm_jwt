package main

import (
	//"gorm.io/driver/postgres"

	connection "golang_gin_gorm_jwt/checkError"
	"golang_gin_gorm_jwt/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	connection.CheckError(err)

	router := gin.Default()

	routes.AuthStudentsRoutes(router)
	routes.AuthAdminRoutes(router)
	router.Run(":8080")

}
