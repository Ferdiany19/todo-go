package controllers

import (
	"net/http"
	"todos-api/src/config"
	"todos-api/src/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB = config.ConnectDB()

type todoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type todoResponse struct {
	todoRequest
	ID uint `json:"id"`
}

func CreateTodo(context *gin.Context) {
	var data todoRequest

	if err := context.ShouldBindJSON(&data); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Request body
	todo := models.Todo{}
	todo.Name = data.Name
	todo.Description = data.Description

	result := db.Create(&todo)
	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "something went wrong"})
		return
	}
	// response nya
	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description

	context.JSON(http.StatusCreated, gin.H{"data": response})
}

func GetAllTodos(context *gin.Context) {
	var todos []models.Todo

	err := db.Find(&todos)
	if err.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Error getting data"})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Success",
		"data":    todos,
	})
}

func GetTodoById(context *gin.Context) {
	var todo models.Todo

	if err := db.Where("id = ?", context.Param("id")).First(&todo).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Error": "record not found!"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": todo})
}

func UpdateTodo(c *gin.Context) {
	var data todoRequest
	reqParamId := c.Param("id")
	// id := cast.ToUint(reqParamId)

	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{}

	query := db.Where("id = ?", reqParamId).First(&todo)
	if query.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data Not Found!"})
		return
	}

	todo.Name = data.Name
	todo.Description = data.Description

	result := db.Save(&todo)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	// ini bagian response
	var response todoResponse
	response.ID = todo.ID
	response.Name = todo.Name
	response.Description = todo.Description
	c.JSON(http.StatusCreated, gin.H{"data": response})
}
