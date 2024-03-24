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

type CreatePhotoRequest struct {
	Title		string	`json:"title"` 
	Caption 	string	`json:"caption"`
	PhotoUrl	string	`json:"photo_url"`
}

var appJson = "application/json"

func CreatePhoto(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	//req body

	var requestBody CreatePhotoRequest
	if err := c.ShouldBindJSON(&requestBody); err != nil {
        helpers.ResponseBadRequest(c, err.Error())
        return
    }

	Photo := models.Photo{}
	userID := uint(userData["id"].(float64))

	Photo.Title = requestBody.Title
	Photo.Caption = requestBody.Caption
	Photo.PhotoUrl = requestBody.PhotoUrl
	Photo.UserID = userID

	if contentType == appJson {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = userID

	err := db.Debug().Create(&Photo).Error

	if err != nil {
		helpers.ResponseError(c, err.Error())
		return
	}

	helpers.ResponseCreated(c, gin.H{
		"message": "Photo successfully created",
		"created": Photo,
	})
}

func GetAllPhoto(c *gin.Context) {
	db := config.GetDB()
	var Photos []models.Photo

	  // Preload only specific fields from the User association
	err := db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
        return db.Select("id, username, email")
    }).Find(&Photos).Error


	if err != nil {
		helpers.ResponseError(c, err.Error())
		return
	}

	helpers.ResponseOK(c, gin.H{
		"message": "All photos successfully to retrieved",
		"data":    Photos,
	})
}

func GetPhotoByID(c *gin.Context) {
	db := config.GetDB()
	var Photos []models.Photo

	photoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResponseBadRequestWithMessage(c, err.Error(), helpers.ID)
		return
	}

	err = db.Debug().First(&Photos, photoId).Error
	if err != nil {
		helpers.ResponseNotFound(c, err.Error())
		return
	}

	helpers.ResponseOK(c, gin.H{
		"message": "Photo successfully retrieved",
		"data":    Photos,
	})
}

func UpdatePhoto(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResponseBadRequestWithMessage(c, err.Error(), helpers.ID)
		return
	}

	if contentType == appJson {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	userID := uint(userData["id"].(float64))
	Photo.UserID = userID
	Photo.ID = uint(photoId)

	err = db.Model(&Photo).Where("id = ?", photoId).Updates(models.Photo{Title: Photo.Title, Caption: Photo.Caption, PhotoUrl: Photo.PhotoUrl}).Error
	if err != nil {
		helpers.ResponseNotFound(c, err.Error())
		return
	}

	helpers.ResponseOK(c, gin.H{
		"updated": Photo,
	})
}

func DeletePhoto(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	photoId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResponseBadRequest(c, err.Error())
		return
	}

	Photo := models.Photo{}

	err = db.Debug().Where("id = ?", photoId).Where("user_id = ?", userID).First(&Photo).Error
	if err != nil {
		helpers.ResponseBadRequest(c, err.Error())
		return
	}

	err = db.Debug().Delete(&Photo).Error
	if err != nil {
		helpers.ResponseBadRequest(c, err.Error())
		return
	}

	helpers.ResponseOK(c, gin.H{
		"message": "Photo has been successfully to deleted",
	})
}
