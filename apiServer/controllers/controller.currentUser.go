package controllers

import (
	"net/http"

	"github.com/Alfazal007/ctr_solana/helpers"
	"github.com/Alfazal007/ctr_solana/internal/database"
	typeconvertor "github.com/Alfazal007/ctr_solana/typeConvertor"
)

func (apiCfg *ApiConf) CurrentUser(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(database.User)
	if !ok {
		helpers.RespondWithError(w, 400, "Login to use this feature")
		return
	}
	helpers.RespondWithJSON(w, 200, typeconvertor.UserConvertor(user))
}
