package controllers

import (
	"fga-final-project-mygram/config"
	"fga-final-project-mygram/helpers"
	"fga-final-project-mygram/models"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateComment(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		helpers.ResponseBadRequestWithMessage(c, err.Error(), helpers.ID)
		return
	}

	Comment := models.Comment{}
	userId := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = userId
	Comment.PhotoID = uint(photoId)

	err = db.Debug().Create(&Comment).Error
	if err != nil {
		helpers.ResponseError(c, err.Error())
		return
	}

	helpers.ResponseCreated(c, gin.H{
		"message": "Comment created successfully",
		"created": Comment,
	})
}

func GetAllComment(c *gin.Context) {
	db := config.GetDB()
	photoId, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		helpers.ResponseBadRequestWithMessage(c, err.Error(), helpers.ID)
		return
	}

	var Comments []models.Comment


	err = db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
        return db.Select("id, username, email")
    }).Where("photo_id = ?", photoId).Find(&Comments).Error


	if err != nil {
		helpers.ResponseError(c, err.Error())
		return
	}

	helpers.ResponseOK(c, gin.H{
		"message":  "Comments retrieved successfully",
		"comments": Comments,
	})
}

func GetCommentById(c *gin.Context) {
	db := config.GetDB()
	Comment := models.Comment{}

	commentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResponseBadRequestWithMessage(c, err.Error(), helpers.ID)
		return
	}

	err = db.Debug().Preload("User").Where("id = ?", commentId).First(&Comment).Error
	if err != nil {
		helpers.ResponseBadRequest(c, err.Error())
		return
	}

	helpers.ResponseOK(c, gin.H{
		"message": "Comment retrieved successfully",
		"Comment": Comment,
	})
}

func UpdateComment(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Comment := models.Comment{}

	commentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResponseBadRequestWithMessage(c, err.Error(), helpers.ID)
		return
	}

	if contentType == appJson {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	userID := uint(userData["id"].(float64))
	Comment.UserID = userID
	Comment.ID = uint(commentId)

	err = db.Model(&Comment).Where("id = ?", commentId).Updates(models.Comment{PhotoID: Comment.PhotoID, Message: Comment.Message}).Error
	if err != nil {
		helpers.ResponseNotFound(c, err.Error())
		return
	}

	helpers.ResponseOK(c, gin.H{
		"updated": gin.H{
			"id":      Comment.ID,
			"user_id": Comment.UserID,
			"message": Comment.Message,
		},
	})
}

func DeletedComment(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	commentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResponseBadRequestWithMessage(c, err.Error(), helpers.ID)
		return
	}

	Comment := models.Comment{}

	err = db.Debug().Where("id = ?", commentId).Where("user_id = ?", userID).First(&Comment).Delete(&Comment).Error
	if err != nil {
		helpers.ResponseNotFound(c, err.Error())
		return
	}

	helpers.ResponseOK(c, gin.H{
		"message": "Photo has been successfully to deleted",
	})
}
