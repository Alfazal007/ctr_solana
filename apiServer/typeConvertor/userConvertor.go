package typeconvertor

import "github.com/Alfazal007/ctr_solana/internal/database"

type UserToBeReturned struct {
	Username string `json:"username"`
	Id       string `json:"id"`
	UserType string `json:"userType"`
}

func UserConvertor(user database.User) UserToBeReturned {
	return UserToBeReturned{
		Username: user.Username,
		Id:       user.ID.String(),
		UserType: string(user.Role.UserRole),
	}
}
