package handlers

import (
	"net/http"

	"github.com/erixalv/crud-golang/gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type TaskHandler struct {
	db *gorm.DB
}

func NewTaskHandler(db *gorm.DB) *TaskHandler {
	return &TaskHandler{db: db}
}

func (h *TaskHandler) ReadTasks(c *gin.Context) {
	var tasks []models.Tasks

	result := h.db.Find(&tasks)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
	}

	c.JSON(http.StatusOK, tasks)
}

func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task models.Tasks
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := h.db.Create(&task)
	if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

	c.JSON(http.StatusCreated, task)
}

func (h *TaskHandler) UpdateTask(c *gin.Context) {
	var task models.Tasks
	id := c.Param("id")

	if h.db.First(&task, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error" : "Task not found."})
		return
	}

	//ShouldBindJson -> sobrescreve os campos da struct task antiga
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.db.Save(&task)
	c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) DeleteTask(c *gin.Context) {
	var task models.Tasks
	id := c.Param("id")

	if h.db.First(&task, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error" : "Task not found."})
		return
	}

	h.db.Delete(&task)
	c.JSON(http.StatusOK, gin.H{"message" : "Task deleted."})
}