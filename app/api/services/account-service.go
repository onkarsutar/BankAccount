package services

import (
	"errors"

	"github.com/onkarsutar/BankAccount/cmd/dbstore"
	"github.com/onkarsutar/BankAccount/model"
)

func GetAllAccountsService() (map[int64]model.Account, error) {
	db := dbstore.GetDB()
	if (*db).Records == nil {
		return nil, errors.New("DB not initialised.")
	}
	(*db).Lock.RLock()
	defer (*db).Lock.RUnlock()

	return (*db).Records, nil
}
func GetAccountService(accountNumber int64) (model.Account, error) {
	db := dbstore.GetDB()
	if (*db).Records == nil {
		return model.Account{}, errors.New("DB not initialised.")
	}
	(*db).Lock.RLock()
	defer (*db).Lock.RUnlock()

	accountObj, ok := (*db).Records[accountNumber]
	if !ok {
		return model.Account{}, errors.New("Account not exists.")
	}
	return accountObj, nil
}
func GetBalanceService(accountNumber int64) (float64, error) {
	db := dbstore.GetDB()
	if (*db).Records == nil {
		return 0, errors.New("DB not initialised.")
	}
	(*db).Lock.RLock()
	defer (*db).Lock.RUnlock()
	if accountObj, ok := (*db).Records[accountNumber]; ok {
		return accountObj.Balance, nil
	}
	return 0, errors.New("Account not exists.")
}

func TransferService(transferObj model.Transfer) error {
	db := dbstore.GetDB()
	(*db).Lock.Lock()
	defer (*db).Lock.Unlock()
	if (*db).Records == nil {
		return errors.New("DB not initialised.")
	}
	if transferObj.CreditAccountNumber == transferObj.DebitAccountNumber {
		return errors.New("Transfer not allowed.")
	}
	creditAccountObj, ok := (*db).Records[transferObj.CreditAccountNumber]
	if !ok {
		return errors.New("Credit account not exists.")
	}
	if transferObj.TransferAmount > creditAccountObj.Balance {
		return errors.New("Insufficient balance.")
	}
	debitAccountObj, ok := (*db).Records[transferObj.DebitAccountNumber]
	if !ok {
		return errors.New("Debit account not exists.")
	}

	creditAccountObj.Balance = creditAccountObj.Balance - transferObj.TransferAmount
	(*db).Records[creditAccountObj.AccountNumber] = creditAccountObj
	debitAccountObj.Balance = debitAccountObj.Balance + transferObj.TransferAmount
	(*db).Records[debitAccountObj.AccountNumber] = debitAccountObj

	return nil
}
