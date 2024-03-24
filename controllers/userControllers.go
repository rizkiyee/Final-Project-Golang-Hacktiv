package controllers

import (
	"fga-final-project-mygram/config"
	"fga-final-project-mygram/helpers"
	"fga-final-project-mygram/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	db := config.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		helpers.ResponseError(c, err.Error())
		return
	}

	helpers.ResponseCreated(c, gin.H{
		"id":       User.ID,
		"email":    User.Email,
		"username": User.Username,
		"age":      User.Age,
	})
}

func UserLogin(c *gin.Context) {
	db := config.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}
	password := ""

	if contentType == appJson {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password

	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error
	if err != nil {
		helpers.ResponseStatusUnauthorizedWithMessage(c, helpers.InvalidUser)
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		helpers.ResponseStatusUnauthorizedWithMessage(c, helpers.InvalidUser)
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email)

	helpers.ResponseOK(c, gin.H{
		"token":   token,
		"message": "User has been successfully to login",
	})
}

func UserDelete(c *gin.Context) {
    db := config.GetDB()
    userIDParam := c.Param("id")
    userID, err := strconv.Atoi(userIDParam)
    if err != nil {
        helpers.ResponseBadRequestWithMessage(c, "Invalid user ID", helpers.ID)
        return
    }

    var user models.User
    if err := db.First(&user, userID).Error; err != nil {
        helpers.ResponseNotFound(c, "User not found")
        return
    }

    if err := db.Delete(&user).Error; err != nil {
        helpers.ResponseError(c, err.Error())
        return
    }
    helpers.ResponseOK(c, gin.H{
        "message": "User deleted successfully",
    })
}

func UserUpdate(c *gin.Context) {
    db := config.GetDB()
    userIDParam := c.Param("id")
    userID, err := strconv.Atoi(userIDParam)
    if err != nil {
        helpers.ResponseBadRequestWithMessage(c, "Invalid user ID", helpers.ID)
        return
    }

    // Retrieve the user from the database
    var user models.User
    if err := db.First(&user, userID).Error; err != nil {
        helpers.ResponseNotFound(c, "User not found")
        return
    }

    // Create a struct to hold the fields to update
    type UserUpdateInput struct {
        Username string `json:"username"`
        Email    string `json:"email"`
    }

    var userInput UserUpdateInput
    if err := c.ShouldBindJSON(&userInput); err != nil {
        helpers.ResponseBadRequestWithMessage(c, "Invalid input data", err.Error())
        return
    }

    // Update user fields
    user.Username = userInput.Username
    user.Email = userInput.Email

    // Validate updated user
    if err := db.Save(&user).Error; err != nil {
        helpers.ResponseError(c, err.Error())
        return
    }

    helpers.ResponseOK(c, gin.H{
        "message": "User updated successfully",
        "user":    user,
    })
}
