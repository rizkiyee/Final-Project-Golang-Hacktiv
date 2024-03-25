package router

import (
	"fga-final-project-mygram/controllers"
	"fga-final-project-mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("register", controllers.UserRegister)
		userRouter.POST("login", controllers.UserLogin)
		userRouter.DELETE("/:id", controllers.UserDelete)
		userRouter.PUT("/:id",  controllers.UserUpdate)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", controllers.GetAllPhoto)
		photoRouter.GET("/:id", controllers.GetPhotoByID)
		photoRouter.PUT("/:id", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:id", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.POST("/photo/:photoId", controllers.CreateComment)
		commentRouter.GET("/photo/:photoId", controllers.GetAllComment)
		commentRouter.GET("/:id", controllers.GetCommentById)
		commentRouter.PUT("/:id", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:id", middlewares.CommentAuthorization(), controllers.DeletedComment)
	}

	socmedRouter := r.Group("/socmeds")
	{
		socmedRouter.Use(middlewares.Authentication())
		socmedRouter.POST("/", controllers.CreateSocmed)
		socmedRouter.GET("/", controllers.GetAllSocmed)
		socmedRouter.GET("/:id", controllers.GetSocmedById)
		socmedRouter.PUT("/:id", middlewares.SocmedAuthorization(), controllers.UpdateSocmed)
		socmedRouter.DELETE("/:id", middlewares.SocmedAuthorization(), controllers.DeleteSocmed)
	}

	return r
}
