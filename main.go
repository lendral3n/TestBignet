package main

import (
	"tesBignet/auth"
	"tesBignet/config"
	"tesBignet/db"
	midd "tesBignet/middleware"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	cfg := config.InitConfig()

	DB := db.InitDB(cfg)

	authRepo := auth.NewRepository(DB)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewAuthHandler(authService)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/register", authHandler.Register)
	e.POST("/login", authHandler.Login)
	e.GET("/users", authHandler.GetAllUsers, midd.JWTMiddleware())

	e.Logger.Fatal(e.Start(":8080"))
}
