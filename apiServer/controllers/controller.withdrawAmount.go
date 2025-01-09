package controllers

import (
	"bytes"
	"crypto/ed25519"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Alfazal007/ctr_solana/helpers"
	"github.com/Alfazal007/ctr_solana/internal/database"
	"github.com/Alfazal007/ctr_solana/utils"
)

func (apiCfg *ApiConf) Withdraw(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(database.User)
	if !ok {
		helpers.RespondWithError(w, 400, "Invalid user")
		return
	}

	if user.Role.UserRole != "labeller" {
		helpers.RespondWithError(w, 400, "Only for labellers")
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

	message := user.ID.String()
	messageBytes := []byte(message)

	valid := ed25519.Verify(ed25519.PublicKey(pk), messageBytes, signature)

	if !valid {
		helpers.RespondWithError(w, 400, "Invalid signature")
		return
	}

	labeller, err := apiCfg.DB.GetLabellerBalance(r.Context(), user.ID)
	if err != nil {
		helpers.RespondWithError(w, 400, "No funds to withdraw")
		return
	}
	newLabellerBalanceInt, err := strconv.ParseInt(labeller.Lamports, 10, 64)
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue converting string balance to int")
		return
	}
	tx, err := apiCfg.SQLDB.BeginTx(r.Context(), &sql.TxOptions{})
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue creating the transaction")
		return
	}
	defer tx.Rollback()
	txContextQueuries := apiCfg.DB.WithTx(tx)
	err = txContextQueuries.DeductBalance(r.Context(), database.DeductBalanceParams{Lamports: "0", LabellerID: user.ID})
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue updating the balance")
		return
	}
	_, err = makePostRequest(WithdrawParams{
		Secret:   utils.LoadEnvVariables().ApiSecret,
		Lamports: newLabellerBalanceInt,
		To:       requestBody.PublicKey,
	})
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue withdrawing the amount")
		return
	}
	err = tx.Commit()
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue committing the transaction")
		return
	}

	helpers.RespondWithJSON(w, 200, []string{})
}

type WithdrawParams struct {
	Secret   string `json:"secret"`
	To       string `json:"to"`
	Lamports int64  `json:"lamports"`
}

func makePostRequest(payload WithdrawParams) (string, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error marshalling payload: %v", err)
	}
	url := "http://localhost:8002/transfer"

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making POST request: %v", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	// Check the status code
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("received non-OK HTTP status: %d, response: %s", resp.StatusCode, string(body))
	}

	return string(body), nil
}
