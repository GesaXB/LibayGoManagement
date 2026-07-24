package controllers

import (
	"net/http"

	requestdto "github.com/GesaXB/LibayGoManagement/dto/requestDto"
	responsedto "github.com/GesaXB/LibayGoManagement/dto/responseDto"
	"github.com/GesaXB/LibayGoManagement/services"
	"github.com/gin-gonic/gin"
)

type authController struct {
	service services.AuthService
}

func NewAuthController(service services.AuthService) *authController {
	return &authController{
		service: service,
	}
}

func (c *authController) Register(ctx *gin.Context) {

	var req requestdto.RegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, err := c.service.Register(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Register success",
		"token":   token,
	})
}

func (c *authController) Login(ctx *gin.Context) {

	var req requestdto.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	token, user, err := c.service.Login(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, responsedto.AuthResponse{
		Token: token,
		User: responsedto.UserResponse{
			ID:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		},
	})
}

func (c *authController) Me(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "token valid",
	})
}
