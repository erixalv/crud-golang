package handlers

import (
	"net/http"

	"github.com/erixalv/crud-golang/gin/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) ReadUsers(c *gin.Context) {
	var users []models.User

	result := h.db.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	result := h.db.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if h.db.First(&user, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error" : "User not found."})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	result := h.db.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")

	if h.db.First(&user, id).Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error" : "User not found."})
		return
	}

	result := h.db.Delete(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error" : result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message" : "Task deleted."})
}