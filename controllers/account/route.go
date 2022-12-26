package account

import (
	"github.com/labstack/echo"
)

func Route(g *echo.Echo) {

	e := g.Group("/account")

	e.GET("/list", GetAccounts)
	e.GET("/:id", GetAccount)
	e.POST("/", CreateAccount)
	e.PATCH("/:id", UpdateAccount)
	e.DELETE("/:id", DeleteAccount)

}
