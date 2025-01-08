package controllers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Alfazal007/ctr_solana/helpers"
	"github.com/Alfazal007/ctr_solana/internal/database"
	typeconvertor "github.com/Alfazal007/ctr_solana/typeConvertor"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *ApiConf) GetProjectToVote(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(database.User)
	if !ok {
		helpers.RespondWithError(w, 400, "Login again")
		return
	}
	if user.Role.UserRole != "labeller" {
		helpers.RespondWithError(w, 400, "Only for labellers")
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
		helpers.RespondWithError(w, 400, "Issue finding the project")
		return
	}

	if project.Completed.Bool {
		helpers.RespondWithError(w, 400, "Project already complete")
		return
	}

	if !project.Started.Bool {
		helpers.RespondWithError(w, 400, "Project not started yet")
		return
	}

	_, err = apiCfg.DB.GetExistingVote(r.Context(), database.GetExistingVoteParams{
		VoterID:   user.ID,
		ProjectID: project.ID,
	})

	if err == sql.ErrNoRows {
		fmt.Println("Here")
		projectImages, err := apiCfg.DB.GetProjectImages(r.Context(), project.ID)
		if err != nil {
			helpers.RespondWithError(w, 400, "Issue finding the project data")
			return
		}
		helpers.RespondWithJSON(w, 200, typeconvertor.ProjectImageData(projectImages))
		return
	}

	if err != nil {
		helpers.RespondWithError(w, 400, "Issue finding the project")
		return
	}
	helpers.RespondWithError(w, 400, "Already voted")
}
