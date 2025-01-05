package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Alfazal007/ctr_solana/helpers"
	"github.com/Alfazal007/ctr_solana/internal/database"
	"github.com/google/uuid"
)

type RequestBodyStruct struct {
	Secret   string `json:"secret"`
	Address  string `json:"address"`
	Lamports string `json:"lamports"`
}

func (apiCfg *ApiConf) IncreaseBalance(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBodyStruct
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	fmt.Println(requestBody)
	if err != nil || requestBody.Address == "" || requestBody.Lamports == "" || requestBody.Secret == "" {
		helpers.RespondWithError(w, 400, "")
		return
	}
	creatorId, err := uuid.Parse(requestBody.Address)
	if err != nil {
		helpers.RespondWithError(w, 400, "Invalid address")
		return
	}
	tx, err := apiCfg.SQLDB.BeginTx(r.Context(), &sql.TxOptions{})
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue creating the transaction")
		return
	}
	defer tx.Rollback()
	txContextQueuries := apiCfg.DB.WithTx(tx)
	creatorBalance, err := txContextQueuries.GetCreatorBalance(r.Context(), creatorId)
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue talking to the database")
		return
	}
	newBalance, err := strconv.ParseInt(requestBody.Lamports, 10, 64)
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue converting the to be added balance from string")
		return
	}
	balanceOfCreatorInI64, err := strconv.ParseInt(creatorBalance.Lamports, 10, 64)
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue converting the balance from string")
		return
	}
	balanceOfCreatorInI64 = balanceOfCreatorInI64 + newBalance
	newCreatorBalance := strconv.FormatInt(balanceOfCreatorInI64, 10)
	txContextQueuries.DeductCreatorBalance(r.Context(), database.DeductCreatorBalanceParams{
		Lamports:  newCreatorBalance,
		CreatorID: creatorBalance.CreatorID,
	})
	err = tx.Commit()
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue committing the transaction")
		return
	}
	helpers.RespondWithJSON(w, 200, []string{})
}
