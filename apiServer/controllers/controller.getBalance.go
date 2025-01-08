package controllers

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"

	"github.com/Alfazal007/ctr_solana/helpers"
	"github.com/Alfazal007/ctr_solana/internal/database"
)

func (apiCfg *ApiConf) FetchBalance(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(database.User)
	if !ok {
		helpers.RespondWithError(w, 400, "Login to use this feature")
		return
	}
	var balanceAmount string
	if user.Role.UserRole == "creator" {
		userBalance, err := apiCfg.DB.GetCreatorBalance(r.Context(), user.ID)
		if err == sql.ErrNoRows {
			helpers.RespondWithError(w, 400, "You dont have a registered key yet")
			return
		}
		if err != nil {
			helpers.RespondWithError(w, 400, "Issue talking to the database")
			return
		}
		balanceAmount = userBalance.Lamports
	} else {
		userBalance, err := apiCfg.DB.GetLabellerBalance(r.Context(), user.ID)
		if err == sql.ErrNoRows {
			helpers.RespondWithError(w, 400, "You dont have a registered key yet")
			return
		}
		if err != nil {
			helpers.RespondWithError(w, 400, "Issue talking to the database")
			return
		}
		balanceAmount = userBalance.Lamports
	}
	balanceInFloat, err := strconv.ParseFloat(strings.TrimSpace(balanceAmount), 64)
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue converting the balance to float")
		return
	}
	balanceInFloat = balanceInFloat / 1000000000
	helpers.RespondWithJSON(w, 200, balanceInFloat)
}
