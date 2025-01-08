package controllers

import (
	"database/sql"
	"net/http"

	"github.com/Alfazal007/ctr_solana/helpers"
	"github.com/Alfazal007/ctr_solana/internal/database"
)

type ResponsePK struct {
	IsRegistered bool   `json:"isRegistered"`
	PublicKey    string `json:"publicKey"`
}

func (apiCfg *ApiConf) GetPublicKey(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(database.User)
	if !ok {
		helpers.RespondWithError(w, 400, "Issue fetching user")
		return
	}

	creatorPublicKey, err := apiCfg.DB.GetCreatorBalance(r.Context(), user.ID)
	if err == sql.ErrNoRows {
		helpers.RespondWithJSON(w, 200, ResponsePK{
			IsRegistered: false,
			PublicKey:    "",
		})
		return
	}
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue fetching from the database")
		return
	}
	helpers.RespondWithJSON(w, 200, ResponsePK{
		IsRegistered: true,
		PublicKey:    creatorPublicKey.CreatorPkBs64.String,
	})
}
