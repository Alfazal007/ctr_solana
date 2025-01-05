package controllers

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Alfazal007/ctr_solana/helpers"
	"github.com/Alfazal007/ctr_solana/internal/database"
	"github.com/google/uuid"
)

type RequstBody struct {
	Signature string `json:"signature"`
	PublicKey string `json:"publicKey"`
}

func (apiCfg *ApiConf) AddCreatorPK(w http.ResponseWriter, r *http.Request) {
	// Parse string to UUID (assuming this is hardcoded)
	uuidStr := "550e8400-e29b-41d4-a716-446655440000"
	id, _ := uuid.Parse(uuidStr)
	creator := database.User{ID: id, Username: "someone", Password: "smoe"}

	// Decode the request body
	var requestBody RequstBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil || requestBody.PublicKey == "" || requestBody.Signature == "" {
		helpers.RespondWithError(w, 400, "Invalid request body")
		return
	}

	// Decode the signature from Base64
	signature, err := base64.StdEncoding.DecodeString(requestBody.Signature)
	if err != nil {
		log.Fatal("Failed to decode Base64 signature:", err)
	}

	// Decode the public key from Base64
	pk, _ := base64.StdEncoding.DecodeString(requestBody.PublicKey)

	// Hash the UUID string with the "SOLANA" prefix
	message := "SOLANA" + creator.ID.String()
	hash := sha256.New()
	hash.Write([]byte(message))
	hashedMessage := hash.Sum(nil)

	// Verify the signature using Ed25519
	valid := ed25519.Verify(ed25519.PublicKey(pk), hashedMessage, signature)
	if !valid {
		fmt.Println("Wrong")
		helpers.RespondWithError(w, 400, "Invalid signature")
		return
	}

	// If valid, respond with a success message
	fmt.Println("Correct")
	helpers.RespondWithJSON(w, 200, []string{})
}
