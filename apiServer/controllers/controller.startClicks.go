package controllers

import (
	"net/http"

	"github.com/Alfazal007/ctr_solana/helpers"
	"github.com/Alfazal007/ctr_solana/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *ApiConf) StartVote(w http.ResponseWriter, r *http.Request) {
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
	if project.Started.Bool {
		helpers.RespondWithJSON(w, 200, "Already started")
		return
	}
	_, err = apiCfg.DB.StartProject(r.Context(), project.ID)
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue starting the project")
		return
	}
	helpers.RespondWithJSON(w, 200, []string{})
}
