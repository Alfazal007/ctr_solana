package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Alfazal007/ctr_solana/helpers"
	"github.com/Alfazal007/ctr_solana/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

type PublicIdType struct {
	PublicId string `json:"publicId"`
}

func (apiCfg *ApiConf) CreateVote(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(database.User)
	if !ok {
		helpers.RespondWithError(w, 400, "Login to use this feature")
		return
	}
	if user.Role.UserRole == "creator" {
		helpers.RespondWithError(w, 400, "Only labellers can do this")
		return
	}
	// check if project started and not ended
	projectId := chi.URLParam(r, "projectId")
	projectIdInUUid, err := uuid.Parse(projectId)
	if err != nil {
		helpers.RespondWithError(w, 400, "Invalid project id")
		return
	}
	var publicIdStruct PublicIdType
	err = json.NewDecoder(r.Body).Decode(&publicIdStruct)
	if err != nil || publicIdStruct.PublicId == "" {
		helpers.RespondWithError(w, 400, "Invalid request body")
		return
	}
	publicId := publicIdStruct.PublicId
	project, err := apiCfg.DB.GetExistingProjectById(r.Context(), projectIdInUUid)
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue finding the project in the database")
		return
	}
	if !project.Started.Bool || project.Completed.Bool {
		helpers.RespondWithError(w, 400, "Project ended or not started yet")
		return
	}

	// check number of votes < 50
	if project.Votes >= 50 {
		helpers.RespondWithError(w, 400, "Maximum votes done")
		return
	}

	_, err = apiCfg.DB.GetImageByPublicId(r.Context(), database.GetImageByPublicIdParams{
		PublicID:  publicId,
		ProjectID: project.ID,
	})
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue talking to the database while finding the image")
		return
	}

	// check if already voted
	_, err = apiCfg.DB.GetExistingVote(r.Context(), database.GetExistingVoteParams{
		VoterID:   user.ID,
		ProjectID: project.ID,
	})

	if err != nil && err != sql.ErrNoRows {
		helpers.RespondWithError(w, 400, "Issue talking to the database")
		return
	}

	if err != sql.ErrNoRows {
		helpers.RespondWithError(w, 400, "Cannot vote 2 times for the same project")
		return
	}

	// update balance of labeller and creator and vote as well
	tx, err := apiCfg.SQLDB.BeginTx(r.Context(), &sql.TxOptions{})
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue creating the transaction")
		return
	}
	defer tx.Rollback()
	txContextQueuries := apiCfg.DB.WithTx(tx)
	// check creator balance and decrement it by 0.1 lamports
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
	newBalance := balanceOfCreatorInI64 - 10000000
	newCreatorBalance := strconv.FormatInt(newBalance, 10)
	err = txContextQueuries.DeductCreatorBalance(r.Context(), database.DeductCreatorBalanceParams{
		Lamports:  newCreatorBalance,
		CreatorID: project.CreatorID,
	})
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue deducting the balance of creator")
		return
	}

	// check receiver balance and increment it by 0.1 lamports
	labellerBalance, err := txContextQueuries.GetLabellerBalance(r.Context(), user.ID)
	if err != nil && err != sql.ErrNoRows {
		helpers.RespondWithError(w, 400, "Issue getting labeller balance from database")
		return
	}
	var newLabellerBalance int64
	if err == sql.ErrNoRows {
		newLabellerBalance = 10000000
	} else {
		newLabellerBalance, err = strconv.ParseInt(labellerBalance.Lamports, 10, 64)
		if err != nil {
			helpers.RespondWithError(w, 400, "Issue converting the labeller balance from string")
			return
		}
		newLabellerBalance = newLabellerBalance + 10000000
	}

	newLabellerBalanceString := strconv.FormatInt(newLabellerBalance, 10)
	err = txContextQueuries.UpsertLabellerBalance(r.Context(), database.UpsertLabellerBalanceParams{
		LabellerID: user.ID,
		Lamports:   newLabellerBalanceString,
	})
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue upserting the balance of labeller")
		return
	}

	// balances have been updated, now to create the vote
	_, err = txContextQueuries.CreateVote(r.Context(), database.CreateVoteParams{
		VoterID:   user.ID,
		ProjectID: project.ID,
		PublicID:  publicId,
	})
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue creating the vote")
		return
	}
	// update project table to have one more vote
	projectInTx, err := txContextQueuries.GetExistingProjectById(r.Context(), project.ID)
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue finding the project within the transaction")
		return
	}
	newVoteNumber := projectInTx.Votes + 1
	err = txContextQueuries.IncreaseVoteCount(r.Context(), database.IncreaseVoteCountParams{
		ID:    project.ID,
		Votes: newVoteNumber,
	})
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue increasing the vote count")
		return
	}
	err = tx.Commit()
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue committing the transaction")
		return
	}
	helpers.RespondWithJSON(w, 200, []string{})
}
