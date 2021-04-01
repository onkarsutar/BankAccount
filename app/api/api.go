package api

import (
	"net/http"

	"github.com/labstack/echo"
	echoMiddleware "github.com/labstack/echo/middleware"
	"github.com/onkarsutar/BankAccount/app/api/controllers"
	"github.com/onkarsutar/BankAccount/app/middleware"
	"github.com/onkarsutar/BankAccount/model"
)

// Init : Initialise routes
func Init(e *echo.Echo) {
	o := e.Group("/o")
	r := e.Group("/r")

	r.Use(echoMiddleware.JWT([]byte(model.JWTKey)))

	middleware.Init(e, o, r)
	controllers.Init(o, r)

	e.GET("/", RootRoute)
	e.GET("/checkststus", CheckStatusRoute)
}

// RootRoute : Base route of Endpoint
func RootRoute(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome")
}

// CheckStatusRoute : Route to check endpoint status
func CheckStatusRoute(c echo.Context) error {
	return c.String(http.StatusOK, "Running..")
}
