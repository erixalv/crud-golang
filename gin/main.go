package main

import (
	"github.com/erixalv/crud-golang/gin/config"
	"github.com/erixalv/crud-golang/gin/models"
	"github.com/erixalv/crud-golang/gin/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db := config.SetupDatabase()
	db.AutoMigrate(&models.Tasks{})

	r := gin.Default()
	routes.SetupRoutes(r, db)

	r.Run("127.0.0.1:8000")
}