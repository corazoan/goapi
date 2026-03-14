package handler

import (
	"encoding/json"
	"net/http"

	"github.com/corazoan/api/api"
	"github.com/corazoan/api/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
)

func GetAccountDetails(w http.ResponseWriter, r *http.Request) {
	var params = api.AccountDetailsParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	if err = decoder.Decode(&params, r.URL.Query()); err != nil {
		api.InternalErrorHandler(w, err)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(w, err)
		return
	}

	var accountDetails = (*database).GetAccountDetails(params.Username)

	if accountDetails == nil {
		log.Error(err)
		api.InternalErrorHandler(w, err)
		return
	}

	var response = api.AccountDetailsResponse{
		Username:  params.Username,
		Balance:   accountDetails.Balance,
		AccountId: accountDetails.AccountId,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w, err)
		return
	}

}
