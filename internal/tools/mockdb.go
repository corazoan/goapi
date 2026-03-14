package tools

import (
	"time"
)

type mockDb struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username:  "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username:  "marie",
	},
}

var mockAccountDetails = map[string]AccountDetails{
	"alex": {
		Balance:   2000,
		AccountId: 1234,
	},
	"jason": {
		Balance:   4000,
		AccountId: 5678,
	},
	"marie": {
		Balance:   6000,
		AccountId: 91011,
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"jason": {
		Coins:    200,
		Username: "jason",
	},
	"marie": {
		Coins:    300,
		Username: "marie",
	},
}

func (d *mockDb) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second + 1)
	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDb) GetAccountDetails(username string) *AccountDetails {
	time.Sleep(time.Second + 1)
	var clientData = AccountDetails{}
	clientData, ok := mockAccountDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDb) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second + 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &clientData
}

func (d *mockDb) SetUpDatabase() error {
	return nil
}
