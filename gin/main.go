package main

import (
	"github.com/erixalv/crud-golang/gin/config"
	"github.com/erixalv/crud-golang/gin/models"
	"github.com/erixalv/crud-golang/gin/routes"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	db := config.SetupDatabase()
	db.AutoMigrate(&models.Tasks{})

	r := gin.Default()
	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
        AllowHeaders:     []string{"Content-Type", "Authorization"},
    }))
	routes.SetupRoutes(r, db)

	r.Run("127.0.0.1:8000")
}