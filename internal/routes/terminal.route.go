package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/handarudwiki/mkp-tht/config"
	"github.com/handarudwiki/mkp-tht/internal/controllers"
	"github.com/handarudwiki/mkp-tht/internal/middlewares"
	"github.com/handarudwiki/mkp-tht/internal/repositories"
	"github.com/handarudwiki/mkp-tht/internal/services"
	"gorm.io/gorm"
)

func InitTerminal(db *gorm.DB, jwt config.JWT, router *gin.Engine) {
	terminalRepository := repositories.NewTerminalRepository(db)
	terminalService := services.NewTerminalService(terminalRepository)
	terminalController := controllers.NewTerminalController(terminalService)

	terminal := router.Group("/terminal")
	terminal.Use(middlewares.AuthMiddleware(jwt))
	{
		terminal.POST("/", terminalController.Create)
	}
}
