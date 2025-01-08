package typeconvertor

import (
	"github.com/Alfazal007/ctr_solana/internal/database"
)

type ProjectImage struct {
	SecureUrl string `json:"secureUrl"`
	PublicID  string `json:"publicId"`
}

func ProjectImageData(projects []database.ProjectImage) []ProjectImage {
	projectsToBeReturned := make([]ProjectImage, 0)
	for i := 0; i < len(projects); i++ {
		projectsToBeReturned = append(projectsToBeReturned, ProjectImage{
			SecureUrl: projects[i].SecureUrl,
			PublicID:  projects[i].PublicID,
		})
	}
	return projectsToBeReturned
}
