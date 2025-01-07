package controllers

import (
	"net/http"

	"github.com/Alfazal007/ctr_solana/helpers"
	"github.com/Alfazal007/ctr_solana/internal/database"
	typeconvertor "github.com/Alfazal007/ctr_solana/typeConvertor"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *ApiConf) VotedProjects(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(database.User)
	if !ok {
		helpers.RespondWithError(w, 400, "Login again")
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
		helpers.RespondWithError(w, 400, "Issue fetching data from database")
		return
	}
	if project.CreatorID != user.ID {
		helpers.RespondWithError(w, 400, "You are not the creator")
		return
	}
	if !project.Completed.Bool {
		helpers.RespondWithError(w, 400, "The project is not completed yet")
		return
	}
	votes_and_public_id, err := apiCfg.DB.GetVotesForProject(r.Context(), project.ID)
	helpers.RespondWithJSON(w, 200, typeconvertor.VotesMany(votes_and_public_id))
}
