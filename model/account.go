package model

import "time"

// AccountJSON : Entity that holds Account Details from JSON
type AccountJSON struct {
	ID      string
	Name    string
	Balance string
}

// Account : Entity that holds Account Details from DB
type Account struct {
	ID            string
	AccountNumber int64
	Name          string
	Balance       float64
	CreatedOn     *time.Time
	ModifiedOn    *time.Time
}

// Transfer : Entity that holds Balance Transfer Details
type Transfer struct {
	CreditAccountNumber int64   `json:"creditAccountNumber"`
	DebitAccountNumber  int64   `json:"debitAccountNumber"`
	TransferAmount      float64 `json:"transferAmount"`
}
