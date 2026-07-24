package controllers

import (
	"net/http"
	"strconv"

	requestdto "github.com/GesaXB/LibayGoManagement/dto/requestDto"
	"github.com/GesaXB/LibayGoManagement/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func (c categoryController) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 2)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	category, err := c.service.GetById(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "category not found",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, category)
}

func (c categoryController) Create(ctx *gin.Context) {
	var req requestdto.CategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	category, err := c.service.Create(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, category)
}

func (c categoryController) Update(ctx *gin.Context) {
	var req requestdto.UpdateCategoryRequest

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	updateCategory, err := c.service.Update(uint(id), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, updateCategory)
}
