package v1

import (
	"encoding/json"

	res "../../response"
	serV1 "../../services/v1"
	"github.com/gin-gonic/gin"
)

func UserList(c *gin.Context) {
	var userService serV1.UserService

	//decode the request body into struct and failed if any error occur
	err := json.NewDecoder(c.Request.Body).Decode(&userService.User)
	if err != nil {
		res.Respond(c.Writer, res.Message(1, "Invalid request"))
		return
	}

	//call service
	resp := userService.UserList()

	//return response using api helper
	res.Respond(c.Writer, resp)

}
