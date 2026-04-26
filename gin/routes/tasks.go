package routes

import (
	"github.com/erixalv/crud-golang/gin/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TaskRoutes(r *gin.Engine, db *gorm.DB) {
	taskHandler := handlers.NewTaskHandler(db)

	tasks := r.Group("/tasks") 
	{
		tasks.GET("", taskHandler.ReadTasks)
		tasks.POST("", taskHandler.CreateTask)
		tasks.PUT("/:id", taskHandler.UpdateTask)
		tasks.DELETE("/:id", taskHandler.DeleteTask)
	}
}