package controllers

import (
	"net/http"

	"github.com/GesaXB/LibayGoManagement/services"
	"github.com/gin-gonic/gin"
)

type categoryController struct {
	service services.CategoryService
}

func NewCategoryController(s services.CategoryService) *categoryController {
	return &categoryController{s}
}

func (c categoryController) GetAll(ctx *gin.Context) {
	categories, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, categories)
}
