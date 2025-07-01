package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/handarudwiki/mkp-tht/config"
	"github.com/handarudwiki/mkp-tht/internal/controllers"
	"github.com/handarudwiki/mkp-tht/internal/repositories"
	"github.com/handarudwiki/mkp-tht/internal/services"
	"gorm.io/gorm"
)

func InitUser(db *gorm.DB, jwt config.JWT, router *gin.Engine) {
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository, jwt.Secret)
	userController := controllers.NewUserController(userService)

	user := router.Group("/user")
	{
		user.POST("/register", userController.Register)
		user.POST("/login", userController.Login)
	}
}
