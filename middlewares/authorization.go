package middlewares

import (
	"fga-final-project-mygram/config"
	"fga-final-project-mygram/helpers"
	"fga-final-project-mygram/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)


func UserAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.GetDB()
		userId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			helpers.ResponseBadRequestWithMessage(c, err.Error(), helpers.ID)
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		User := models.User{}

		err = db.Select("user_id").First(&User, uint(userId)).Error
		if err != nil {
			helpers.ResponseNotFound(c, err.Error())
			return
		}

		if User.ID != userID {
			helpers.ResponseStatusUnauthorizedWithMessage(c, "Not allowed to access this data")
			return
		}

		c.Next()
	}
}


func PhotoAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.GetDB()
		photoId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			helpers.ResponseBadRequestWithMessage(c, err.Error(), helpers.ID)
			return
		}

		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Photo := models.Photo{}

		err = db.Select("user_id").First(&Photo, uint(photoId)).Error
		if err != nil {
			helpers.ResponseNotFound(c, err.Error())
			return
		}

		if Photo.UserID != userID {
			helpers.ResponseStatusUnauthorizedWithMessage(c, "Not allowed to access this data")
			return
		}

		c.Next()
	}
}

func CommentAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.GetDB()
		commentId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			helpers.ResponseBadRequestWithMessage(c, err.Error(), helpers.ID)
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Comment := models.Comment{}

		err = db.Select("user_id").First(&Comment, uint(commentId)).Error
		if err != nil {
			helpers.ResponseNotFound(c, err.Error())
			return
		}

		if Comment.UserID != userID {
			helpers.ResponseStatusUnauthorizedWithMessage(c, helpers.NotAllowedAccessData)
			return
		}

		c.Next()
	}
}

func SocmedAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := config.GetDB()
		socmedId, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			helpers.ResponseBadRequestWithMessage(c, err.Error(), helpers.ID)
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Socmed := models.SocialMedia{}

		err = db.Select("user_id").First(&Socmed, uint(socmedId)).Error
		if err != nil {
			helpers.ResponseNotFound(c, err.Error())
			return
		}

		if Socmed.UserID != userID {
			helpers.ResponseStatusUnauthorizedWithMessage(c, helpers.NotAllowedAccessData)
			return
		}

		c.Next()
	}
}
