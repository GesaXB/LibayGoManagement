package controllers

import (
	"net/http"

	requestdto "github.com/GesaXB/LibayGoManagement/dto/requestDto"
	"github.com/GesaXB/LibayGoManagement/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type authorController struct {
	service services.AuthorService
}

func NewAuthorController(s services.AuthorService) *authorController {
	return &authorController{
		service: s,
	}
}

func (c authorController) GetAll(ctx *gin.Context) {
	auhtors, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, auhtors)
}

func (c authorController) GetById(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	author, err := c.service.GetById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	ctx.JSON(http.StatusOK, author)
}

func (c *authorController) Create(ctx *gin.Context) {
	var inputAuthor requestdto.AuthorRequest
	err := ctx.ShouldBindJSON(&inputAuthor)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	newAuthor, err := c.service.Create(inputAuthor)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, newAuthor)
}

func (c *authorController) Update(ctx *gin.Context) {
	var input requestdto.AuthorRequest

	id := ctx.Param("id")
	if _, err := uuid.Parse(id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid id"})
		return
	}

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	authorUpdated, err := c.service.Update(id, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, authorUpdated)

}
