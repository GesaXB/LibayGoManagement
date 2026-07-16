package routes

import (
	"net/http"

	"github.com/GesaXB/LibayGoManagement/controllers"
	"github.com/GesaXB/LibayGoManagement/repositories"
	"github.com/GesaXB/LibayGoManagement/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {

	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authController := controllers.NewAuthController(authService)

	userService := services.NewUserService(&userRepo)
	userController := controllers.NewUserController(userService)

	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryController := controllers.NewCategoryController(categoryService)

	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	api := r.Group("/api")
	{
		api.GET("/users", userController.GetAll)
		api.GET("/user/:id", userController.GetById)

		api.GET("/categories", categoryController.GetAll)
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World!",
		})
	})
}
