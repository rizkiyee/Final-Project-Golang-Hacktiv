package controllers

import (
	"errors"
	"fga-final-project-mygram/config"
	"fga-final-project-mygram/helpers"
	"fga-final-project-mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateSocmed(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Socmed := models.SocialMedia{}

	userID := uint(userData["id"].(float64))

	if contentType == appJson {
		c.ShouldBindJSON(&Socmed)
	} else {
		c.ShouldBind(&Socmed)
	}

	Socmed.UserID = userID

	if err := db.Debug().First(&Socmed, "user_id = ?", Socmed.UserID).Scan(&Socmed).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = db.Debug().Create(&Socmed).Scan(&Socmed).Error
			if err != nil {
				helpers.ResponseBadRequest(c, err.Error())
				return
			}
		} else {
			if err != nil {
				helpers.ResponseNotFound(c, err.Error())
				return
			}
		}
	}

	helpers.ResponseCreated(c, gin.H{
		"message": "Social media has been successfully to created",
		"created": Socmed,
	})
}

func GetAllSocmed(c *gin.Context) {
	db := config.GetDB()
	var Socialmedias []models.SocialMedia

	err := db.Debug().Preload("User", func(db *gorm.DB) *gorm.DB {
        return db.Select("id, username, email")
    }).Find(&Socialmedias).Error
	
	if err != nil {
		helpers.ResponseBadRequest(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"socmed": Socialmedias,
	})
}

func GetSocmedById(c *gin.Context) {
	db := config.GetDB()
	var Socialmedias []models.SocialMedia

	socmedId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResponseBadRequestWithMessage(c, err.Error(), helpers.ID)
		return
	}

	err = db.Debug().Preload("User").Where("id = ?", socmedId).First(&Socialmedias).Error
	if err != nil {
		helpers.ResponseNotFound(c, err.Error())
		return
	}

	helpers.ResponseOK(c, gin.H{
		"message": "Photo successfully retrieved",
		"socmed":  Socialmedias,
	})
}

func UpdateSocmed(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Socmed := models.SocialMedia{}

	socmedId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		helpers.ResponseBadRequestWithMessage(c, err.Error(), helpers.ID)
		return
	}

	if contentType == appJson {
		c.ShouldBindJSON(&Socmed)
	} else {
		c.ShouldBind(&Socmed)
	}

	userID := uint(userData["id"].(float64))
	Socmed.UserID = userID
	Socmed.ID = uint(socmedId)

	err = db.Model(&Socmed).Where("id = ?", socmedId).Updates(models.SocialMedia{Name: Socmed.Name, SocialMediaUrl: Socmed.SocialMediaUrl}).Error
	if err != nil {
		helpers.ResponseNotFound(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"updated": Socmed,
	})
}

func DeleteSocmed(c *gin.Context) {
	db := config.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))
	socmedId, _ := strconv.Atoi(c.Param("id"))
	Socmed := models.SocialMedia{}

	err := db.Debug().Where("id = ?", socmedId).Where("user_id = ?", userID).First(&Socmed).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad request",
			"message": err.Error(),
		})
		return
	}

	err = db.Debug().Delete(&Socmed).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err":     "failed to delete photo",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "your photo has been succesfully deleted",
	})
}
