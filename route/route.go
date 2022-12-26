package route

import (
	"fmt"
	"net/http"

	account_router "go_test/controllers/account"
	user_router "go_test/controllers/user"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		message := fmt.Sprintf("Welcome to the project")
		return c.String(http.StatusOK, message)
	})

	user_router.Route(e)
	account_router.Route(e)

	// account_router.Route(g.Group("/api/v1/user"))
	// user_router.Route(g.Group("/api/v1/account"))

}
