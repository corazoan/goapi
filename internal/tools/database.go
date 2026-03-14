package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	AuthToken string
	Username  string
}

type CoinDetails struct {
	Coins    int64
	Username string
}

type AccountDetails struct {
	Balance   int64
	AccountId int64
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserCoins(username string) *CoinDetails
	SetUpDatabase() error
	GetAccountDetails(username string) *AccountDetails
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDb{}
	var err error = database.SetUpDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &database, nil
}
