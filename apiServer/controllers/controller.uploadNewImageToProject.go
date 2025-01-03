package controllers

import (
	"net/http"

	"github.com/Alfazal007/ctr_solana/helpers"
	"github.com/Alfazal007/ctr_solana/internal/database"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

const maxImageSize = 60 * 1024

type SuccessMessage struct {
	Message string `json:"message"`
}

func (apiCfg *ApiConf) GetUrlToUploadImage(w http.ResponseWriter, r *http.Request) {
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
		helpers.RespondWithError(w, 400, "Issue finding the project")
		return
	}
	if user.ID != project.CreatorID.UUID {
		helpers.RespondWithError(w, 400, "You are not the creator")
		return
	}
	if project.Completed.Bool {
		helpers.RespondWithError(w, 400, "Create a new project as this is already complete")
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, maxImageSize)
	err = r.ParseMultipartForm(maxImageSize)
	if err != nil {
		helpers.RespondWithError(w, 400, "File should be under 60kb")
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()
	cloudinaryRes, err := apiCfg.Cloudinary.Upload.Upload(r.Context(), file, uploader.UploadParams{
		Folder:       "web3ctr/" + user.ID.String() + "/" + projectId + "/",
		ResourceType: "image",
	})
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue uploading to cloudinary")
		return
	}
	_, err = apiCfg.DB.CreateProjectImage(r.Context(), database.CreateProjectImageParams{
		PublicID:  cloudinaryRes.PublicID,
		ProjectID: project.ID,
		SecureUrl: cloudinaryRes.SecureURL,
	})
	if err != nil {
		helpers.RespondWithError(w, 400, "Issue saving state in the database")
		return
	}
	helpers.RespondWithJSON(w, 200, SuccessMessage{Message: "Uploaded image successfully"})
}
