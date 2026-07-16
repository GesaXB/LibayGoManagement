package controllers

import (
	"net/http"
	"strconv"

	"github.com/GesaXB/LibayGoManagement/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	services services.UserService
}

func NewUserController(c services.UserService) *UserController {
	return &UserController{c}
}

func (c UserController) GetAll(ctx *gin.Context) {
	users, err := c.services.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) GetById(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.ParseUint(idParam, 10, 2)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid id",
		})
		return
	}

	user, err := c.services.GetById(uint(id))
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.JSON(http.StatusNotFound, gin.H{
				"message": "user not found",
			})
			return
		}
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
