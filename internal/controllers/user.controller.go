package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/handarudwiki/mkp-tht/internal/dto"
	"github.com/handarudwiki/mkp-tht/internal/services"
	"github.com/handarudwiki/mkp-tht/internal/utils"
)

type (
	User interface {
		Register(ctx *gin.Context)
		Login(ctx *gin.Context)
	}
	userController struct {
		service services.User
	}
)

func NewUserController(service services.User) User {
	return &userController{
		service: service,
	}
}

func (c *userController) Register(ctx *gin.Context) {
	var dto dto.Register
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		utils.ResponseError(ctx, "Invalid input", http.StatusBadRequest)
		return
	}

	errors := utils.ValidateRequest(dto)

	if len(errors) > 0 {
		utils.ResponseValidationError(ctx, errors)
		return
	}

	user, err := c.service.Register(ctx.Request.Context(), dto)
	if err != nil {
		utils.ResponseError(ctx, err.Error(), utils.GetHttpStatusCode(err))
		return
	}

	utils.ResponseSuccess(ctx, user)
}

func (c *userController) Login(ctx *gin.Context) {
	var dto dto.Register
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		utils.ResponseError(ctx, "Invalid input", http.StatusBadRequest)
		return
	}

	errors := utils.ValidateRequest(dto)

	if len(errors) > 0 {
		utils.ResponseValidationError(ctx, errors)
		return
	}

	token, err := c.service.Login(ctx.Request.Context(), dto)
	if err != nil {
		utils.ResponseError(ctx, err.Error(), utils.GetHttpStatusCode(err))
		return
	}

	utils.ResponseSuccess(ctx, gin.H{"token": token})
}
