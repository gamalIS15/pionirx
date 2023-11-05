package routes

import (
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
	"net/http"
)

// Handler route
func index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}
func test(c echo.Context) error {
	return c.String(http.StatusOK, "Test")
}

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", index)
	e.GET("/test", test)
	return e
}
