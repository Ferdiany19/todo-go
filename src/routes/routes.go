package routes

import (
	"todos-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	route := gin.Default()

	route.POST("/todo", controllers.CreateTodo)
	route.GET("/todos", controllers.GetAllTodos)
	route.GET("/todos/:id", controllers.GetTodoById)
	route.PUT("/todos/:id", controllers.UpdateTodo)
	// route.DELETE("/todo/:idTodo", controllers.DeleteTodo)

	route.Run()
}
