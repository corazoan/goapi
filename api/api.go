package api

import (
	"encoding/json"
	"net/http"
)

type AccountDetailsParams struct {
	Username string
}

type AccountDetailsResponse struct {
	Username  string
	Balance   int64
	AccountId int64
}

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	Code int
	//response code
	Balance int64
}

type Error struct {
	Code    int
	Message string
}

func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, "An unexpected Error Occured.", http.StatusInternalServerError)
	}
)
