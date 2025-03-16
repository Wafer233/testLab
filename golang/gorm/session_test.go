package gorm

import "testing"

func TestSessions(t *testing.T) {
	db := InitDB()
	Sessions(db)
}

func TestSessionsV2(t *testing.T) {
	db := InitDB()
	SessionsV2(db)
}

func TestTransactions(t *testing.T) {
	db := InitDB()
	Transactions(db)
}

func TestTransactionsV2(t *testing.T) {
	db := InitDB()
	TransactionsV2(db)
}
