package typeconvertor

import "github.com/Alfazal007/ctr_solana/internal/database"

type GetVotesForProjectRowReturn struct {
	PublicID  string `json:"publicId"`
	VoteCount int64  `json:"voteCount"`
	SecureUrl string `json:"secureUrl"`
}

func VotesMany(projects []database.GetVotesForProjectRow) []GetVotesForProjectRowReturn {
	projectsToBeReturned := make([]GetVotesForProjectRowReturn, 0)
	for i := 0; i < len(projects); i++ {
		projectsToBeReturned = append(projectsToBeReturned, GetVotesForProjectRowReturn{
			PublicID:  projects[i].PublicID,
			VoteCount: projects[i].VoteCount,
			SecureUrl: projects[i].SecureUrl,
		})
	}
	return projectsToBeReturned
}
