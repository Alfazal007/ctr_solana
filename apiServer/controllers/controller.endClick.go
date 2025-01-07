package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/Alfazal007/ctr_solana/helpers"
	"github.com/Alfazal007/ctr_solana/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *ApiConf) EndVote(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(database.User)
	if !ok {
		helpers.RespondWithError(w, 400, "Login to use this feature")
		return
	}
	if user.Role.UserRole != "creator" {
		helpers.RespondWithError(w, 400, "Only creators can do this")
		return
	}

	projectId := chi.URLParam(r, "projectId")
	projectIdInUUid, err := uuid.Parse(projectId)
	if err != nil {
		helpers.RespondWithError(w, 400, "Invalid project id")
		return
	}
	project, err := apiCfg.DB.GetExistingProjectById(r.Context(), projectIdInUUid)
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue finding the project in the database")
		return
	}
	if user.ID != project.CreatorID {
		helpers.RespondWithError(w, 400, "You are not the creator")
		return
	}
	if !project.Started.Bool {
		helpers.RespondWithError(w, 400, "Not started yet")
		return
	}
	if project.Completed.Bool {
		helpers.RespondWithError(w, 400, "Already terminated")
		return
	}
	// =====================
	tx, err := apiCfg.SQLDB.BeginTx(r.Context(), &sql.TxOptions{})
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue creating the transaction")
		return
	}
	defer tx.Rollback()
	txContextQueuries := apiCfg.DB.WithTx(tx)
	creatorBalance, err := txContextQueuries.GetCreatorBalance(r.Context(), project.CreatorID)
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue getting creator balance")
		return
	}
	balanceOfCreatorInI64, err := strconv.ParseInt(creatorBalance.Lamports, 10, 64)
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue converting the balance from string")
		return
	}
	newBalance := balanceOfCreatorInI64 - 100000000
	newCreatorBalance := strconv.FormatInt(newBalance, 10)
	err = txContextQueuries.DeductCreatorBalance(r.Context(), database.DeductCreatorBalanceParams{
		Lamports:  newCreatorBalance,
		CreatorID: project.CreatorID,
	})
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue deducting the balance of creator")
		return
	}
	_, err = txContextQueuries.EndProject(r.Context(), project.ID)
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue starting the project")
		return
	}
	err = tx.Commit()
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue committing the transaction")
		return
	}
	helpers.RespondWithJSON(w, 200, []string{})
}
