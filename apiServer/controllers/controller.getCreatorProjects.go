package controllers

import (
	"database/sql"
	"net/http"

	"github.com/Alfazal007/ctr_solana/helpers"
	"github.com/Alfazal007/ctr_solana/internal/database"
	typeconvertor "github.com/Alfazal007/ctr_solana/typeConvertor"
)

func (apiCfg *ApiConf) GetCreatorProjects(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(database.User)
	if !ok {
		helpers.RespondWithError(w, 400, "Issue getting the user")
		return
	}
	if user.Role.UserRole != "creator" {
		helpers.RespondWithError(w, 400, "This user is not a creator")
		return
	}
	projects, err := apiCfg.DB.GetCreatorProjects(r.Context(), user.ID)
	if err != nil && err != sql.ErrNoRows {
		helpers.RespondWithError(w, 400, "Issue talking to the database")
		return
	}
	if err == sql.ErrNoRows {
		helpers.RespondWithJSON(w, 200, []string{})
		return
	}
	helpers.RespondWithJSON(w, 200, typeconvertor.ProjectConvertorForCreatorMany(projects))
}
