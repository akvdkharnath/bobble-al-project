package user

import "github.com/labstack/echo"

func Route(e *echo.Echo) {

	g := e.Group("/user")

	g.GET("/list", GetUsers)
	g.GET("/:id", GetUser)
	g.POST("/", CreateUser)
	g.PATCH("/:id", UpdateUser)
	g.DELETE("/:id", DeleteUser)
}
