package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/onkarsutar/BankAccount/app/api/services"
	"github.com/onkarsutar/BankAccount/helper/loghelper"
	"github.com/onkarsutar/BankAccount/model"
)

// Init : Initialise Handler related to Account
func Init(o, r *echo.Group) {
	o.GET("/api/getallaccounts", GetAllAccountsHandler)
	o.GET("/api/getbyid/:accountNumber", GetAccountHandler)
	o.GET("/api/getbalance/:accountNumber", GetBalanceHandler)
	o.POST("/api/transfer", TransferHandler)
}

// GetAllAccountsHandler : Handler function to get all account details
func GetAllAccountsHandler(c echo.Context) error {
	db, err := services.GetAllAccountsService()
	if err != nil {
		loghelper.LogError("GetAllAccountsHandler Error: ", err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}
	return c.JSON(http.StatusOK, db)
}

// GetAccountHandler : Handler function to get account details
func GetAccountHandler(c echo.Context) error {
	accountNumber, _ := strconv.ParseInt(c.Param("accountNumber"), 10, 64)
	accountObj, err := services.GetAccountService(accountNumber)
	if err != nil {
		loghelper.LogError("GetAccountHandler Error: ", err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}

	return c.JSON(http.StatusOK, accountObj)
}

// GetBalanceHandler : Handler function to get balance
func GetBalanceHandler(c echo.Context) error {
	accountNumber, _ := strconv.ParseInt(c.Param("accountNumber"), 10, 64)
	accountBalance, err := services.GetBalanceService(accountNumber)
	if err != nil {
		loghelper.LogError("GetBalanceHandler Error: ", err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}

	return c.JSON(http.StatusOK, accountBalance)
}

// TransferHandler : Handler function to handle balance transfer
func TransferHandler(c echo.Context) error {
	transferObj := model.Transfer{}
	err := c.Bind(&transferObj)
	if err != nil {
		loghelper.LogError("TransferHandler Error: ", err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}

	err = services.TransferService(transferObj)
	if err != nil {
		loghelper.LogError("GetBalanceHandler Error: ", err)
		return c.JSON(http.StatusExpectationFailed, err.Error())
	}
	return c.JSON(http.StatusOK, "Success")
}
