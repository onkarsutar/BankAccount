package main

import (
	"strconv"

	"github.com/onkarsutar/BankAccount/app"
	"github.com/onkarsutar/BankAccount/cmd/dbstore"
	"github.com/onkarsutar/BankAccount/helper/confighelper"
	"github.com/onkarsutar/BankAccount/helper/loghelper"
)

func main() {
	Init()
	app.StartServer()

}

// Init : Initialise required dependencies
func Init() {
	confighelper.InitViper()
	dbstore.InitData()

	maxBackupCount, _ := strconv.Atoi(confighelper.GetConfig("maxBackupCount"))
	maxBackupFileSize, _ := strconv.Atoi(confighelper.GetConfig("maxBackupFileSize"))
	maxAgeForBackupFiles, _ := strconv.Atoi(confighelper.GetConfig("maxAgeForBackupFiles"))

	loghelper.Init("./logs/server.log", true, maxBackupCount, maxBackupFileSize, maxAgeForBackupFiles, true)

}
