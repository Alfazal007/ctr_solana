package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/Alfazal007/ctr_solana/helpers"
	"github.com/Alfazal007/ctr_solana/internal/database"
	typeconvertor "github.com/Alfazal007/ctr_solana/typeConvertor"
	"github.com/Alfazal007/ctr_solana/utils"
	"github.com/google/uuid"
)

type CreateUserBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	UserType string `json:"userType"`
}

func (apiCfg *ApiConf) CreateUser(w http.ResponseWriter, r *http.Request) {
	var createUserBody CreateUserBody
	err := json.NewDecoder(r.Body).Decode(&createUserBody)
	if err != nil || createUserBody.Username == "" || createUserBody.Password == "" {
		helpers.RespondWithError(w, 400, "Invalid request body format, provide both username and password")
		return
	}
	if len(createUserBody.Username) > 20 {
		helpers.RespondWithError(w, 400, "Length of the username should be between 1 and 20(inclusive)")
		return
	}

	if len(createUserBody.Password) <= 7 || len(createUserBody.Password) > 20 {
		helpers.RespondWithError(w, 400, "Length of the password should be between 8 and 20(inclusive)")
		return
	}

	existingUser, err := apiCfg.DB.CheckSimilarUserExists(r.Context(), createUserBody.Username)
	if err != nil && err != sql.ErrNoRows {
		helpers.RespondWithError(w, 400, "Issue talking to the database")
		return
	}
	if existingUser > 0 {
		helpers.RespondWithError(w, 400, "Use different username")
		return
	}
	hashedPassword, err := utils.HashPassword(createUserBody.Password)
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue hashing the password")
		return
	}

	var userRole database.NullUserRole
	if createUserBody.UserType == "creator" {
		userRole = database.NullUserRole{
			Valid:    true,
			UserRole: database.UserRoleCreator,
		}
	} else {
		userRole = database.NullUserRole{
			Valid:    true,
			UserRole: database.UserRoleLabeller,
		}
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:       uuid.New(),
		Username: createUserBody.Username,
		Password: hashedPassword,
		Role:     userRole,
	})
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue talking to the database")
		return
	}

	helpers.RespondWithJSON(w, 201, typeconvertor.UserConvertor(user))
}
