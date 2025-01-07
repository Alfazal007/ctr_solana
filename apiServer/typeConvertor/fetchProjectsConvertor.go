package typeconvertor

import (
	"github.com/Alfazal007/ctr_solana/internal/database"
)

type FetchProjectsToVoteRowToSend struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ProjectToVote(projects []database.FetchProjectsToVoteRow) []FetchProjectsToVoteRowToSend {
	projectsToBeReturned := make([]FetchProjectsToVoteRowToSend, 0)
	for i := 0; i < len(projects); i++ {
		projectsToBeReturned = append(projectsToBeReturned, FetchProjectsToVoteRowToSend{
			Name: projects[i].Name,
			ID:   projects[i].ID.String(),
		})
	}
	return projectsToBeReturned
}
