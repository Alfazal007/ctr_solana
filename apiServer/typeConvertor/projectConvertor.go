package typeconvertor

import "github.com/Alfazal007/ctr_solana/internal/database"

type ProjectToBeReturned struct {
	Name      string `json:"name"`
	Id        string `json:"id"`
	CreatorId string `json:"creatorId"`
}

func ProjectConvertor(project database.Project) ProjectToBeReturned {
	return ProjectToBeReturned{
		Name:      project.Name,
		Id:        project.ID.String(),
		CreatorId: project.CreatorID.UUID.String(),
	}
}
