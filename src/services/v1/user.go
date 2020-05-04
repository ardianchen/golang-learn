package v1service

import (
	"../../models"
	res "../../response"
)

type UserService struct {
	User models.User
}

func (us *UserService) UserList() map[string]interface{} {
	// os.Stdout.WriteString(us.User.email)
	user := &models.User{}

	userData := models.User{
		ID:    1,    // user.ID,
		Name:  "ar", //user.Name,
		Email: "as", // user.Email,
	}
	userDatas := result.UserResponse{
		ID:    1,    // user.ID,
		Name:  "ar", //user.Name,
		Email: "as", // user.Email,
	}
	response := res.Message(0, "This is from version 1 api")
	response["data"] = userDatas
	return response
}
