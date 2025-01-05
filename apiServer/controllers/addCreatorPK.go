package controllers

import (
	"net/http"

	"github.com/Alfazal007/ctr_solana/helpers"
	"github.com/Alfazal007/ctr_solana/internal/database"
)

func (apiCfg *ApiConf) AddCreatorPK(w http.ResponseWriter, r *http.Request) {
	creator, ok := r.Context().Value("user").(database.User)
	if !ok {
		helpers.RespondWithError(w, 400, "Issue finding the user")
		return
	}
}
