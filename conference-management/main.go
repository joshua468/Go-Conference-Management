package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joshua468/conference-management/routes"
	"github.com/joshua468/conference-management/storage"
)

func main() {
	db := storage.InitDB()
	defer db.Close()

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")
}
