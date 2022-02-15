package main

import (
	"fmt"
	"yuno/echo-example/controller"
	"yuno/echo-example/database"
	_ "yuno/echo-example/docs"
	middleWare "yuno/echo-example/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
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
	router.Use(middleWare.MiddlewareAuth)

	// Routes
	router.GET("/student", controller.GetStudents)
	router.POST("/student", controller.CreateStudent)
	router.DELETE("/student/:student", controller.DeleteStudent)
	router.PATCH("/student/:student", controller.UpdateStudent)
	router.GET("/health-check", controller.HealthCheck)
	router.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	router.Logger.Fatal(router.Start(":8080"))
}
