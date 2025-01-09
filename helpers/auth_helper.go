package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUserType(c *gin.Context, role string) (err error) {
	userType := c.GetString("user_type")

	err = nil
	if userType != role {
		err = errors.New("unauthorized")
		return err
	}
	return err
}

func MatchUserTypeToUid(ctx *gin.Context, user_id string) (err error) {
	userType := ctx.GetString("user_type")
	userId := ctx.GetString("user_id")
	err = nil

	if userType == "USER" && userId != user_id {
		err = errors.New("unauthorized")
		return err
	}

	err = CheckUserType(ctx, userType)
	return err
}
