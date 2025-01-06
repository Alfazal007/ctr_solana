package controllers

import (
	"net/http"

	"github.com/Alfazal007/ctr_solana/helpers"
)

func (apiCfg *ApiConf) Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "accessToken",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteNoneMode,
		MaxAge:   -1,
	}
	http.SetCookie(w, &cookie)
	helpers.RespondWithJSON(w, 200, "Logout successful")
}
