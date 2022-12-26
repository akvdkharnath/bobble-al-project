package main

import (
	"fmt"
	"go_test/db"
	"go_test/route"

	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	cwd, _ := os.Getwd()
	godotenv.Load(cwd + `/.env`)
	PORT := os.Getenv("APP_PORT")
	fmt.Println("PORT", PORT)
	db.Init()

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))
	route.Init(e)
	e.Logger.Fatal(e.Start(":" + PORT))

}
