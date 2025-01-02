package controllers

import (
	"net/http"

	"github.com/Alfazal007/ctr_solana/helpers"
)

type DummyMessage struct {
	Message string `json:"message"`
}

func (apicCfg *ApiConf) CreateUser(w http.ResponseWriter, r *http.Request) {
	helpers.RespondWithJSON(w, 200, DummyMessage{Message: "hello world"})
}
