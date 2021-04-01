package app

import (
	"github.com/labstack/echo"
	"github.com/onkarsutar/BankAccount/app/api"
	"github.com/onkarsutar/BankAccount/helper/confighelper"
)

// StartServer : Starts REST API endpoint
func StartServer() {
	e := echo.New()
	api.Init(e)
	apiServerPort := confighelper.GetConfig("apicdnServerPort")
	e.Start(":" + apiServerPort)
}
