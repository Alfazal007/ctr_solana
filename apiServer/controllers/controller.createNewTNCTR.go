package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Alfazal007/ctr_solana/helpers"
	"github.com/Alfazal007/ctr_solana/internal/database"
	typeconvertor "github.com/Alfazal007/ctr_solana/typeConvertor"
	"github.com/google/uuid"
)

type CreateNewCTRType struct {
	Name string `json:"name"`
}

func (apiCfg *ApiConf) CreateNewCTR(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("user").(database.User)
	if !ok {
		helpers.RespondWithError(w, 400, "Login to use this feature")
		return
	}
	if string(user.Role.UserRole) != "creator" {
		helpers.RespondWithError(w, 400, "User is not a creator")
		return
	}
	var createNewCTRType CreateNewCTRType
	err := json.NewDecoder(r.Body).Decode(&createNewCTRType)
	if err != nil || createNewCTRType.Name == "" || len(createNewCTRType.Name) > 20 {
		helpers.RespondWithError(w, 400, "The project name should be greater than 1 and less than 20(inclusive)")
		return
	}
	// check if the other projects are over or not
	countRunningTasks, err := apiCfg.DB.CountRunningProjects(r.Context(), user.ID)
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue talking to the database")
		return
	}
	if countRunningTasks > 0 {
		helpers.RespondWithError(w, 400, "Complete the previous project to start the new project")
		return
	}

	newProject, err := apiCfg.DB.CreateProject(r.Context(), database.CreateProjectParams{
		ID:        uuid.New(),
		Name:      createNewCTRType.Name,
		CreatorID: user.ID,
	})
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue creating new project, try changing the project name")
		return
	}
	helpers.RespondWithJSON(w, 200, typeconvertor.ProjectConvertor(newProject))
}
