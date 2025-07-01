package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/handarudwiki/mkp-tht/internal/dto"
	"github.com/handarudwiki/mkp-tht/internal/services"
	"github.com/handarudwiki/mkp-tht/internal/utils"
)

type (
	Terminal interface {
		Create(ctx *gin.Context)
	}

	terminalController struct {
		service services.Terminal
	}
)

func NewTerminalController(service services.Terminal) Terminal {
	return &terminalController{
		service: service,
	}
}

func (c *terminalController) Create(ctx *gin.Context) {
	var dto dto.CreateTerminal
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		utils.ResponseError(ctx, "Invalid input", http.StatusBadRequest)
		return
	}

	errors := utils.ValidateRequest(dto)

	if len(errors) > 0 {
		utils.ResponseValidationError(ctx, errors)
		return
	}

	terminal, err := c.service.Create(ctx.Request.Context(), dto)
	if err != nil {
		utils.ResponseError(ctx, err.Error(), utils.GetHttpStatusCode(err))
		return
	}

	utils.ResponseSuccess(ctx, terminal)
}
