package controllers

import (
	"net/http"

	"github.com/Alfazal007/ctr_solana/helpers"
	"github.com/Alfazal007/ctr_solana/internal/database"
	typeconvertor "github.com/Alfazal007/ctr_solana/typeConvertor"
)

func (apiCfg *ApiConf) ProjectsToVote(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(database.User)
	if !ok {
		helpers.RespondWithError(w, 400, "User not logged in")
		return
	}
	if user.Role.UserRole != "labeller" {
		helpers.RespondWithError(w, 400, "Only for labellers")
		return
	}
	projectsToVote, err := apiCfg.DB.FetchProjectsToVote(r.Context(), user.ID)
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue fetching the projects")
		return
	}
	helpers.RespondWithJSON(w, 200, typeconvertor.ProjectToVote(projectsToVote))
}
