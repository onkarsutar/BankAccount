package dbstore

import (
	"sync"

	"github.com/onkarsutar/BankAccount/model"
)

// DB : In memory data store
type Database struct {
	Records map[int64]model.Account
	Lock    sync.RWMutex
}

var (
	store                Database
	once                 sync.Once
	accountNumberCounter int64 = 1000000000
	numOfWorkers               = 10
)

// GetDB : Returns address of Database
func GetDB() *Database {
	once.Do(func() {
		store.Records = make(map[int64]model.Account)
	})
	return &store
}

// InitDB : Init DB map
// func InitDB(done chan struct{}) {
// }
