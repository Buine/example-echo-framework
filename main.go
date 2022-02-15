package main

import (
	"fmt"
	"yuno/echo-example/controller"
	"yuno/echo-example/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	err := database.ConnectPostgres()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// Echo instance
	router := echo.New()

	// Middleware
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	// Routes
	router.GET("/student", controller.GetStudents)
	router.POST("/student", controller.CreateStudent)
	router.DELETE("/student/:student", controller.DeleteStudent)

	// Start server
	router.Logger.Fatal(router.Start(":8080"))
}
