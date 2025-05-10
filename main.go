package main

import (
	"formative-14/configs"
	"formative-14/database/connection"
	"formative-14/database/migration"
	"formative-14/modules/bioskop"

	"github.com/gin-gonic/gin"
)

func main() {
	configs.Initiator()
	connection.Initiator()
	migration.Initiator(connection.DBConnections)
	initiateRouter()
	defer connection.DBConnections.Close()
}

func initiateRouter() {
	router := gin.Default()
	bioskop.Initiator(router)
	router.Run(":8080")
}