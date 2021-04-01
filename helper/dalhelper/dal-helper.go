package dalhelper

import (
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/onkarsutar/GoSocial/helper/confighelper"
)

var sqlConnection *dbr.Connection
var connectionError error
var sqlOnce sync.Once

func GetSQLConnection() (*dbr.Connection, error) {
	sqlOnce.Do(func() {
		connection, err := dbr.Open("mysql", confighelper.GetConfig("mySQLDSN"), nil)
		if err != nil {
			fmt.Println(err)
			connectionError = err
		}
		connection.SetMaxIdleConns(100)
		connection.SetMaxOpenConns(5000)
		duration := 3 * 24 * time.Hour
		connection.SetConnMaxLifetime(duration)
		sqlConnection = connection
	})
	return sqlConnection, connectionError
}
