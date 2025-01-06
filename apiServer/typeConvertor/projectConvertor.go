package typeconvertor

import (
	"github.com/Alfazal007/ctr_solana/internal/database"
)

type ProjectToBeReturned struct {
	Name      string `json:"name"`
	Id        string `json:"id"`
	CreatorId string `json:"creatorId"`
}

func ProjectConvertor(project database.Project) ProjectToBeReturned {
	return ProjectToBeReturned{
		Name:      project.Name,
		Id:        project.ID.String(),
		CreatorId: project.CreatorID.String(),
	}
}

type ProjectToBeReturnedForCreator struct {
	Name      string `json:"name"`
	Id        string `json:"id"`
	CreatorId string `json:"creatorId"`
	Started   bool   `json:"started"`
	Completed bool   `json:"completed"`
}

func ProjectConvertorForCreator(project database.Project) ProjectToBeReturnedForCreator {
	return ProjectToBeReturnedForCreator{
		Name:      project.Name,
		Id:        project.ID.String(),
		CreatorId: project.CreatorID.String(),
		Started:   project.Started.Bool,
		Completed: project.Completed.Bool,
	}
}

func ProjectConvertorForCreatorMany(projects []database.Project) []ProjectToBeReturnedForCreator {
	projectsToBeReturned := make([]ProjectToBeReturnedForCreator, 0)
	for i := 0; i < len(projects); i++ {
		projectsToBeReturned = append(projectsToBeReturned, ProjectToBeReturnedForCreator{
			Name:      projects[i].Name,
			Id:        projects[i].ID.String(),
			CreatorId: projects[i].CreatorID.String(),
			Started:   projects[i].Started.Bool,
			Completed: projects[i].Completed.Bool,
		})
	}
	return projectsToBeReturned
}
