package middlewares

import (
	"fga-final-project-mygram/helpers"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			helpers.ResponseStatusUnauthorizedWithMessage(c, err.Error())
			return
		}

		c.Set("userData", verifyToken)
		c.Next()
	}
}
