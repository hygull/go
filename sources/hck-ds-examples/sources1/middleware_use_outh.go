package main

import "github.com/labstack/echo"
import "github.com/labstack/echo/middleware"
import "net/http"

func main() {
	e := echo.New()
	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Group level middleware
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string) bool {
		if username == "joe" && password == "secret" {
			return true
		}
		return false
	}))

	// Route level middleware
	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to /users")
			return next(c)
		}
	}
	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users")
	}, track)

	e.Logger.Fatal(e.Start(":8080"))
}
