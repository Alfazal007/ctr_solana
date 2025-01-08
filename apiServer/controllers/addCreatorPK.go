package controllers

import (
	"crypto/ed25519"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/Alfazal007/ctr_solana/helpers"
	"github.com/Alfazal007/ctr_solana/internal/database"
)

type RequstBody struct {
	Signature string `json:"signature"`
	PublicKey string `json:"publicKey"`
}

func (apiCfg *ApiConf) AddCreatorPK(w http.ResponseWriter, r *http.Request) {
	creator, ok := r.Context().Value("user").(database.User)
	if !ok {
		helpers.RespondWithError(w, 400, "Invalid user")
		return
	}

	var requestBody RequstBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil || requestBody.PublicKey == "" || requestBody.Signature == "" {
		helpers.RespondWithError(w, 400, "Invalid request body")
		return
	}

	// Decode the signature from Base64
	signature, err := base64.StdEncoding.DecodeString(requestBody.Signature)
	if err != nil {
		helpers.RespondWithError(w, 400, "Invalid signature")
		return
	}

	// Decode the public key from Base64
	pk, err := base64.StdEncoding.DecodeString(requestBody.PublicKey)
	if err != nil {
		helpers.RespondWithError(w, 400, "Invalid public key")
		return
	}

	message := creator.ID.String()
	messageBytes := []byte(message)

	valid := ed25519.Verify(ed25519.PublicKey(pk), messageBytes, signature)

	if !valid {
		helpers.RespondWithError(w, 400, "Invalid signature")
		return
	}

	// insert into database
	err = apiCfg.DB.AddPublicKey(r.Context(), database.AddPublicKeyParams{
		CreatorPkBs64: sql.NullString{
			Valid: true, String: requestBody.PublicKey,
		},
		CreatorID: creator.ID,
	})
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue adding pk to the database")
		return
	}

	helpers.RespondWithJSON(w, 200, []string{})
}
