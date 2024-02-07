package route

import (
	"logictest/app/todo"
	"logictest/database"

	"github.com/labstack/echo/v4"
)

func InitV1(e *echo.Echo, db database.Database) {

	g := e.Group("/api/v1")

	// add middleware
	// g.Use()

	todoRoute := g.Group("/todo")
	todoService := todo.NewService(db)
	todoHandler := todo.NewHandler(todoService)
	todoRoute.GET("/:id", todoHandler.GetTodoByID)
	todoRoute.GET("/", todoHandler.GetTodoList)
}
