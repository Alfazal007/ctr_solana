package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Alfazal007/ctr_solana/helpers"
	"github.com/Alfazal007/ctr_solana/utils"
)

type LoginUserBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ReturnLoginType struct {
	AccessToken string `json:"accessToken"`
	Id          string `json:"id"`
	Username    string `json:"username"`
}

func (apiCfg *ApiConf) LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginUserBody LoginUserBody
	err := json.NewDecoder(r.Body).Decode(&loginUserBody)
	if err != nil || loginUserBody.Username == "" || loginUserBody.Password == "" {
		helpers.RespondWithError(w, 400, "Invalid request body format, provide both username and password")
		return
	}
	if len(loginUserBody.Username) > 20 {
		helpers.RespondWithError(w, 400, "Length of the username should be between 1 and 20(inclusive)")
		return
	}

	if len(loginUserBody.Password) <= 7 || len(loginUserBody.Password) > 20 {
		helpers.RespondWithError(w, 400, "Length of the password should be between 8 and 20(inclusive)")
		return
	}
	requiredUser, err := apiCfg.DB.GetUserByUsername(r.Context(), loginUserBody.Username)
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue finding the user in the database")
		return
	}
	isPasswordCorrect := utils.VerifyPassword(loginUserBody.Password, requiredUser.Password)
	if !isPasswordCorrect {
		helpers.RespondWithError(w, 400, "Incorrect password")
		return
	}
	accessToken, err := utils.GenerateAccessToken(requiredUser.ID.String(), requiredUser.Username)
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue generating access token")
		return
	}
	cookie := http.Cookie{
		Name:     "accessToken",
		Value:    accessToken,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteNoneMode,
	}
	http.SetCookie(w, &cookie)
	helpers.RespondWithJSON(w, 200,
		ReturnLoginType{
			AccessToken: accessToken,
			Id:          requiredUser.ID.String(),
			Username:    requiredUser.Username,
		})
}
